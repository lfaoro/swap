package app

import (
	"fmt"
	"strconv"
	"strings"

	pb "github.com/lfaoro/swap/gen/go/swap/v1"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	ACCENT_COLOR = lipgloss.Color("10")
	BASE_COLOR   = lipgloss.Color("8")
)

type swapUI struct {
	disable bool

	state *AppState
	api   SwapAPI
	cfg   *Config
	cd    *CoinData

	width  int
	height int

	table       SwapTable
	tableFilter textinput.Model

	viewport    viewport.Model
	detailsView string

	// saved addresses from $HOME/.config/swap/config
	savedAddresses []string
	atAddress      int

	help Help
	log  Log

	sp       spinner.Model
	spinning bool

	pg progress.Model
}

type CoinData struct {
	From *pb.Coin
	To   *pb.Coin

	Exchange string
	TradeID  string
	Fixed    bool
	Payment  bool

	Amount  textinput.Model
	Address textinput.Model
}

func NewTSwapUI(cfg *Config, client pb.CoinServiceClient, debug bool) *swapUI {
	m := &swapUI{
		state: &AppState{},
		api:   NewSwapAPI(client),
		cfg:   cfg,
		cd:    newCoinData(),

		table: NewSwapTable(client),

		spinning: true,
		sp:       spinner.New(spinner.WithSpinner(spinner.Dot)),
		pg: progress.New(
			progress.WithWidth(25),
			progress.WithDefaultGradient(),
			progress.WithoutPercentage(),
		),

		help: NewHelp(),
		log:  NewLog(WithDebug(debug)),
	}

	m.state.Init()
	m.state.GoTo(CoinTable)
	// m.state.GoTo(TrxStatus)
	// m.viewport.SetContent(m.setStatusContent(m.debugStatusData()))
	// m.SetSpinning(true)

	m.tableFilter = textinput.New()
	m.tableFilter.Placeholder = "search..."
	m.tableFilter.CharLimit = 10
	m.tableFilter.Cursor.Blink = true

	m.viewport = viewport.New(m.table.Width(), m.table.Height())

	return m
}

func (m *swapUI) Init() tea.Cmd {
	cmds := []tea.Cmd{
		tea.SetWindowTitle("swapcli.com - Freedom of Exchange"),
		tea.EnterAltScreen,
		// tea.EnableMouseAllMotion,
		// tea.EnableMouseCellMotion,
		tea.EnableBracketedPaste,
		tea.EnableReportFocus,
		m.SetSpinning(true),
		m.table.Init(),
		m.table.Focus(),
		m.log.Init(),
	}

	return tea.Batch(cmds...)
}

func (m *swapUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = DEFAULT_WIDTH
		m.height = min(msg.Height, DEFAULT_HEIGHT)
		m.disable = msg.Width < m.width-10
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 10
		return m, tea.Batch(cmds...)

	case tea.BatchMsg:
		AddLog("main: batch msg")
		cmds = append(cmds, m.SetSpinning(true))

	case tea.MouseMsg:
		if msg.Action == tea.MouseActionPress {
			var keyMsg tea.KeyMsg
			switch msg.Button {
			case tea.MouseButtonWheelUp:
				keyMsg = tea.KeyMsg{Type: tea.KeyUp}
			case tea.MouseButtonWheelDown:
				keyMsg = tea.KeyMsg{Type: tea.KeyDown}
			default:
				return m, nil
			}

			if m.state.IsAt(CoinTable) || m.state.IsAt(RateTable) {
				m.table, cmd = m.table.Update(keyMsg)
				cmds = append(cmds, cmd)
			} else if m.state.IsAt(TrxStatus) {
				m.viewport, cmd = m.viewport.Update(keyMsg)
				cmds = append(cmds, cmd)
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		// TODO: pg test
		case "r":
			cmds = append(cmds, m.pg.DecrPercent(0.01))
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			if m.state.Current() >= InputAddress {
				return m, AddError(fmt.Errorf("cannot change swap/pay at this point: press esc"))
			}
			m.cd.Payment = !m.cd.Payment
			m.table = m.table.SetPayment(m.cd.Payment)
			cmds = append(cmds, AddLog("pay choice: %v", m.cd.Payment))
		}

		// todo(leo): refactor this
		if m.tableFilter.Focused() {
			m.tableFilter, cmd = m.tableFilter.Update(msg)
			cmds = append(cmds, cmd)
			m.table, cmd = m.table.Update(CoinFilterMsg{Query: m.tableFilter.Value()})
			cmds = append(cmds, cmd)

			switch msg.String() {
			case "esc":
				cmds = append(cmds, AddLog("filter: esc pressed"))
				m.tableFilter.Blur()
				m.tableFilter.Reset()
				m.table, cmd = m.table.Update(CoinReqRespMsg{Resp: m.table.coinList})
				m.table.Focus()
				cmds = append(cmds, cmd)
			case "enter":
				cmds = append(cmds, AddLog("filter: enter pressed"))
				m.tableFilter.Blur()
				m.tableFilter.Reset()
				m.table.Focus()
			}
			return m, tea.Batch(cmds...)
		}

		// update tables
		if m.state.IsAt(CoinTable) || m.state.IsAt(RateTable) {
			// cmds = append(cmds, AddLog("state: %v, key: %v", m.state.Current(), msg.String()))
			m.table, cmd = m.table.Update(msg)
			cmds = append(cmds, cmd)
		}

		if m.state.IsAt(CoinTable) {
			switch msg.Type {
			case tea.KeyRunes:
				switch msg.Runes[0] {
				case '/':
					return m, m.tableFilter.Focus()
				}
			case tea.KeyEnter, tea.KeyType(tea.MouseActionRelease):
				cmds = append(cmds, AddLog("table: enter"))
				if m.cd.From == nil {
					m.cd.From = m.table.GetCoin()
				} else {
					coin := m.table.GetCoin()
					if m.cd.From != coin {
						m.cd.To = coin
					} else {
						return m, AddError(fmt.Errorf("cannot swap %s to %s", m.cd.From.Name, coin.Name))
					}

					m.state.GoTo(InputAmount)
					cmds = append(cmds, m.cd.Amount.Focus())
				}
			case tea.KeyEsc:
				return m, tea.Batch(
					AddLog("table: esc"),
					m.resetCoinData(),
					ClearDebug(),
					ClearError(),
					m.table.Init(),
				)

			}
			return m, tea.Batch(cmds...)
		}

		if m.state.IsAt(InputAmount) {
			switch msg.String() {
			default:
				cmds = append(cmds, AddLog("state: %v", m.state.Current()))
				m.cd.Amount, cmd = m.cd.Amount.Update(msg)
				cmds = append(cmds, cmd)

				if m.cd.Amount.Err != nil {
					cmds = append(cmds, AddError(m.cd.Amount.Err))
					m.cd.Amount.Reset()
					m.cd.Amount.Err = nil
					if msg.String() != "esc" {
						return m, tea.Batch(cmds...)
					}
				}

			case "esc":
				m.state.GoTo(CoinTable)
				m.cd.Amount.Err = nil
				m.cd.Amount.Blur()
				m.cd.Amount.Reset()
				cmds = append(cmds, m.table.Init())
				cmds = append(cmds, m.table.Focus())
				return m, tea.Batch(cmds...)

			case "enter":
				//-- validation of min/max dex requirements
				fval := parseFloat(m.cd.Amount.Value())
				min := m.cd.From.Minimum
				max := m.cd.From.Maximum
				ticker := m.cd.From.Ticker
				if m.cd.Payment {
					min = m.cd.To.Minimum
					max = m.cd.To.Maximum
					ticker = m.cd.To.Ticker
				}
				if err := validateAmount(fval, min, max, ticker); err != nil {
					cmds = append(cmds, AddError(err))
					return m, tea.Batch(cmds...)
				}
				//-- validation end

				cmds = append(cmds, AddLog("input amount: %v", m.cd.Amount.Value()))
				cmds = append(cmds, m.SetSpinning(true))
				m.table.state = RateTableState // todo(leo): refactor this
				req := SwapRateReqMsg{
					TickerFrom:  m.cd.From.GetTicker(),
					TickerTo:    m.cd.To.GetTicker(),
					NetworkFrom: m.cd.From.GetNetwork(),
					NetworkTo:   m.cd.To.GetNetwork(),
					Payment:     m.cd.Payment,
				}
				switch req.Payment {
				case true:
					req.AmountTo = m.cd.Amount.Value()
				case false:
					req.AmountFrom = m.cd.Amount.Value()
				}
				cmds = append(cmds, m.api.SwapRate(req))
				cmds = append(cmds, AddLog("rate request: %v", req))
				m.cd.Amount.Blur()
				return m, tea.Batch(cmds...)
			}
			return m, tea.Batch(cmds...)
		}

		if m.state.IsAt(RateTable) {
			switch msg.String() {
			case "esc":
				m.table.rateTable.SetRows(nil) // clear rate table
				m.cd.Exchange = ""
				m.table.state = CoinTableState // todo(leo): refactor
				m.state.GoTo(InputAmount)
				cmds = append(cmds, m.table.Init())
				cmds = append(cmds, m.cd.Amount.Focus())

			case "enter":
				m.cd.Exchange = m.table.GetExchange()
				cmds = append(cmds, AddLog("ratetable: exchange: %s", m.cd.Exchange))

				// populate saved addresses
				m.savedAddresses = m.cfg.GetAllAddress(m.cd.To.Ticker, m.cd.To.Network)

				m.state.GoTo(InputAddress)
				cmds = append(cmds, m.cd.Address.Focus())
			}
		}

		if m.state.IsAt(InputAddress) {
			cmds = append(cmds, AddLog("state: %v", m.state.Current()))
			m.cd.Address, cmd = m.cd.Address.Update(msg)
			cmds = append(cmds, cmd)

			if m.cd.Address.Err != nil {
				cmds = append(cmds, AddError(m.cd.Address.Err))
				cmds = append(cmds, AddLog(m.cd.Address.Err.Error()))
				m.cd.Address.Reset()
				m.cd.Address.Err = nil
				if msg.String() != "esc" {
					return m, tea.Batch(cmds...)
				}
			}

			switch msg.String() {
			case "esc":
				cmds = append(cmds, AddLog("inputaddress: esc pressed"))
				m.cd.Address.Blur()
				m.cd.Address.Reset()
				m.cd.Exchange = ""

				m.state.GoTo(RateTable)
				cmds = append(cmds, m.table.Focus())

			case "enter":
				m.cd.Address, cmd = m.cd.Address.Update(msg)
				cmds = append(cmds, cmd)
				if m.cd.Address.Err != nil {
					cmds = append(cmds, AddError(m.cd.Address.Err))
					cmds = append(cmds, AddLog(m.cd.Address.Err.Error()))
					m.cd.Address.Reset()
					m.cd.Address.Err = nil
					if msg.String() != "esc" {
						return m, tea.Batch(cmds...)
					}
				}

				if m.cd.TradeID == "" {
					return m, AddError(fmt.Errorf("missing tradeID :/"))
				}

				req := SwapTradeReq{m.cd}
				cmds = append(cmds, m.api.SwapTrade(req))

			case "up":
				if len(m.savedAddresses) == 0 {
					return m, AddError(fmt.Errorf("no saved addresses yet"))
				}
				if m.atAddress >= len(m.savedAddresses) {
					m.atAddress = 0
				}
				m.cd.Address.SetValue(m.savedAddresses[m.atAddress])
				m.atAddress++
				m.cd.Address, cmd = m.cd.Address.Update(msg)
				cmds = append(cmds, cmd)
			case "down":
				m.cd.Address.Reset()

			case "ctrl+s":
				if err := m.cfg.SaveAddress(m.cd.To.Ticker, m.cd.To.Network, m.cd.Address.Value()); err != nil {
					cmds = append(cmds, AddLog(err.Error()))
					cmds = append(cmds, AddError(err))
					return m, tea.Batch(cmds...)
				}
				m.savedAddresses = m.cfg.GetAllAddress(m.cd.To.Ticker, m.cd.To.Network)
				_address := m.cd.Address.Value()
				m.cd.Address.Reset()
				cmds = append(cmds, AddError(fmt.Errorf("saved address: %s", _address)))
				return m, tea.Batch(cmds...)

			case "ctrl+d":
				if len(m.savedAddresses) == 0 {
					return m, AddError(fmt.Errorf("no saved addresses yet"))
				}
				// delete address
				m.cfg.DeleteAddress(m.cd.To.Ticker, m.cd.To.Network, m.cd.Address.Value())
				cmds = append(cmds, AddLog("deleted address: %s", m.cd.Address.Value()))
				cmds = append(cmds, AddError(fmt.Errorf("address deleted %s", m.cd.Address.Value())))
				m.savedAddresses = m.cfg.GetAllAddress(m.cd.To.Ticker, m.cd.To.Network)
				m.cd.Address.Reset()

				return m, tea.Batch(cmds...)
			}
		}

		if m.state.IsAt(TrxStatus) {
			m.viewport, cmd = m.viewport.Update(msg)
			cmds = append(cmds, cmd)
			switch msg.String() {
			case "esc":
				m.viewport.SetContent("")
				m.api.Close()
				m.state.GoTo(InputAddress)
				cmds = append(cmds, m.cd.Address.Focus())
			}
		}

		cmds = append(cmds, ClearError())

		// NOTE: keep for debugging
		// cmds = append(cmds, AddLog("processing msg: %v", msg))

	case progress.FrameMsg:
		cmds = append(cmds, AddLog("progress: frame msg"))
		progressModel, cmd := m.pg.Update(msg) // shadows cmd
		m.pg = progressModel.(progress.Model)
		cmds = append(cmds, cmd)

	case spinner.TickMsg:
		AddLog("spinner: tick")
		m.sp, cmd = m.sp.Update(msg)
		if m.spinning {
			cmds = append(cmds, cmd)
		}

	case CoinReqRespMsg:
		cmds = append(cmds, AddLog("main: list coins success"))
		m.table, cmd = m.table.Update(msg)
		cmds = append(cmds, cmd)
		cmds = append(cmds, m.SetSpinning(false))

	case SwapRateRespMsg:
		m.cd.TradeID = msg.Resp.GetTradeId()
		m.table.state = RateTableState
		m.table.Focus()
		m.table, cmd = m.table.Update(msg)
		cmds = append(cmds, cmd)
		cmds = append(cmds, AddLog("main: swap rate success"))
		cmds = append(cmds, m.SetSpinning(false))
		m.state.GoTo(RateTable)

	case SwapTradeRespMsg:
		cmds = append(cmds, m.SetSpinning(true))
		cmds = append(cmds, AddLog("main: SwapTradeResponse"))
		cmds = append(cmds, m.api.SwapStatus(SwapStatusReqMsg{TradeID: msg.GetTradeId()}))
		m.state.GoTo(TrxStatus)

	case SwapStatusRespMsg:
		if msg.Status == "finished" {
			cmds = append(cmds, m.SetSpinning(false))
		}
		m.viewport.SetContent(m.setStatusContent(msg))
		cmds = append(cmds,
			AddLog("main: SwapStatusResponse"),
		)

	case ErrorMsg:
		m.log, cmd = m.log.Update(msg)
		cmds = append(cmds, cmd)
		cmds = append(cmds, m.SetSpinning(false))

	case DebugMsg:
		m.log, cmd = m.log.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *swapUI) View() string {
	if m.disable {
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("3")).
			Bold(true).
			Align(lipgloss.Center, lipgloss.Center).
			Border(lipgloss.ThickBorder(), true).
			BorderForeground(lipgloss.Color("3"))
		msg := "Swap needs a larger terminal window to function properly."
		return style.Render(msg)
	}

	mainView := func() string {
		if m.state.IsAt(TrxStatus) {
			m.viewport.Style = lipgloss.Style{}
			titleStyle := func() string {
				b := lipgloss.RoundedBorder()
				b.Right = "├"
				style := lipgloss.NewStyle().
					Border(b).
					Padding(0, 1)
				title := style.Render("Transaction")
				line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
				return lipgloss.JoinHorizontal(
					lipgloss.Center,
					title,
					line,
				)
			}()

			infoStyle := func() string {
				b := lipgloss.RoundedBorder()
				b.Left = "┤"
				style := lipgloss.NewStyle().BorderStyle(b)
				info := style.Render(fmt.Sprintf("%2.f%%", m.viewport.ScrollPercent()*100))
				line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
				return lipgloss.JoinHorizontal(
					lipgloss.Center,
					line,
					info,
				)
			}()

			return titleStyle + "\n" + m.viewport.View() + "\n" + infoStyle
		} else {
			out := m.table.View()
			out += "\n"
			payView := func() string {
				if m.cd.Payment {
					return m.sp.View() + "[PAY]"
				}
				return m.sp.View() + "[SWAP]"
			}()
			payViewStyle := lipgloss.NewStyle().
				Foreground(ACCENT_COLOR).
				Bold(true)
			baseStyle := lipgloss.NewStyle().
				Foreground(BASE_COLOR).
				Bold(true)
			count := m.table.Width() - len(payView) + len("  ")
			line := strings.Repeat("─", count)
			return out + baseStyle.Render("╰"+payViewStyle.Render(payView)+baseStyle.Render(line+"╯"))
		}
	}()
	coinView := func() string {
		var out string
		if m.cd.From != nil {
			out += fmt.Sprintf("%s (%s)", m.cd.From.Name, m.cd.From.Network)
			if m.cd.To != nil {
				out += fmt.Sprintf(" to %s (%s)", m.cd.To.Name, m.cd.To.Network)
			}
		}
		return out
	}()
	amountView := func() string {
		if m.state.Current() < InputAmount {
			return ""
		}
		var out string
		var ticker string
		if m.cd.Payment {
			ticker = m.cd.To.Ticker
		} else {
			ticker = m.cd.From.Ticker
		}
		out += fmt.Sprintf(" %s%s", m.cd.Amount.View(), ticker)
		return out
	}()
	exchangeView := func() string {
		if m.cd.Exchange == "" {
			return ""
		}
		return fmt.Sprintf(" via %s", m.cd.Exchange)
	}()
	addressView := func() string {
		var out string
		if m.state.Current() < InputAddress {
			return ""
		}
		m.cd.Address.Placeholder = fmt.Sprintf("enter %s address", m.cd.To.Name)

		var width = m.table.Width()
		if len(m.cd.Address.Value()) > width {
			width = len(m.cd.Address.Value()) + 2
		}
		style := lipgloss.NewStyle().
			Width(width).
			Padding(0).
			Border(lipgloss.RoundedBorder(), true).
			BorderForeground(lipgloss.Color("10"))
		out += m.help.AddressHelp() + "\n"
		out += style.Render(m.cd.Address.View())
		return out
	}()

	swapBar := lipgloss.JoinHorizontal(
		lipgloss.Left,
		coinView,
		amountView,
		exchangeView,
	)
	swapBar = lipgloss.NewStyle().
		Width(m.table.Width()).
		Padding(1).
		Border(lipgloss.HiddenBorder(), false).
		BorderForeground(lipgloss.Color("10")).
		Render(swapBar)

	vertical := lipgloss.JoinVertical(
		lipgloss.Left,
		func() string {
			if m.tableFilter.Focused() {
				return m.tableFilter.View()
			}
			return m.help.View()
		}(),
		mainView,
		swapBar,
		addressView,

		m.log.View(),
	)

	if m.state.IsAt(TrxStatus) {
		return mainView
	}

	return vertical
}

func (m *swapUI) resetCoinData() tea.Cmd {
	return func() tea.Msg {
		m.cd.From = nil
		m.cd.To = nil

		m.cd.Amount.Reset()
		m.cd.Amount.Blur()

		m.cd.Address.Reset()
		m.cd.Address.Blur()

		m.cd.Exchange = ""
		m.cd.Payment = false
		return AddLog("main: coin data reset")
	}
}

func tiAddress() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "enter your address..."
	ti.CharLimit = 100
	ti.Cursor.Blink = true
	ti.Validate = func(s string) error {
		if s == "" {
			return fmt.Errorf("address is required")
		}
		if strings.ContainsAny(s, "!@#$%^&*()_+-=[]{}|;:,.<>?/~`\"'\\") {
			return fmt.Errorf("address can only contain letters and numbers")
		}
		return nil
	}
	return ti
}

func tiAmount() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = fmt.Sprintf("%.8f ", 0.0)
	ti.CharLimit = 10
	ti.Cursor.Blink = true
	ti.Validate = func(s string) error {
		if s == "" {
			return fmt.Errorf("amount is required")
		}
		_, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return fmt.Errorf("invalid number")
		}
		return nil
	}
	return ti
}

func validateAmount(amount float64, min, max float64, ticker string) error {
	if amount < min {
		return fmt.Errorf("minimum %v %v", min, ticker)
	}
	if amount > max {
		return fmt.Errorf("maximum %v %v", max, ticker)
	}
	return nil
}

func newCoinData() *CoinData {
	cd := &CoinData{
		From:     nil,
		To:       nil,
		Exchange: "",
		TradeID:  "",
		Amount:   tiAmount(),
		Address:  tiAddress(),
	}
	return cd
}

func (m *swapUI) SetSpinning(enabled bool) tea.Cmd {
	m.spinning = enabled
	if enabled {
		return m.sp.Tick
	}
	return nil
}
