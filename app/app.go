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

type SwapUI struct {
	disable bool

	state *AppState
	api   SwapAPI
	cfg   *Config
	cd    *CoinData

	width  int
	height int

	table       SwapTable
	tableFilter textinput.Model

	trxView viewport.Model

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

func NewTSwapUI(cfg *Config, client pb.CoinServiceClient, debug bool) *SwapUI {
	m := &SwapUI{
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
	// m.trxView.SetContent(m.setStatusContent(m.debugStatusData()))
	// m.SetSpinning(true)

	m.tableFilter = textinput.New()
	m.tableFilter.Placeholder = "search..."
	m.tableFilter.CharLimit = 10
	m.tableFilter.Cursor.Blink = true

	m.trxView = viewport.New(m.table.Width(), m.table.Height())

	return m
}

func (m *SwapUI) Init() tea.Cmd {
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

func (m *SwapUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = DEFAULT_WIDTH
		m.height = min(msg.Height, DEFAULT_HEIGHT)
		m.disable = msg.Width < m.width-10
		m.trxView.Width = msg.Width
		m.trxView.Height = msg.Height - 10
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
				m.trxView, cmd = m.trxView.Update(keyMsg)
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
			m.trxView, cmd = m.trxView.Update(msg)
			cmds = append(cmds, cmd)
			switch msg.String() {
			case "esc":
				m.trxView.SetContent("")
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
		m.trxView.SetContent(m.setStatusContent(msg))
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

func (m *SwapUI) resetCoinData() tea.Cmd {
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

func (m *SwapUI) SetSpinning(enabled bool) tea.Cmd {
	m.spinning = enabled
	if enabled {
		return m.sp.Tick
	}
	return nil
}
