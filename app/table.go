package app

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	pb "github.com/lfaoro/swap/gen/go/swap/v1"
)

type TableState int

const (
	CoinTableState TableState = iota
	RateTableState
)

type SwapTable struct {
	state TableState
	api   SwapAPI

	coinTable     table.Model
	coinTableRows []table.Row
	rateTable     table.Model
	tradeTable    table.Model

	coinList []*pb.Coin

	payment bool
}

func NewSwapTable(client pb.CoinServiceClient) SwapTable {
	columns := []table.Column{
		{Title: "Name", Width: 30},
		{Title: "Ticker", Width: 7},
		{Title: "Network", Width: 10},
		{Title: "Minimum", Width: 10},
		{Title: "Maximum", Width: 15},
	}

	coinTable := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(DEFAULT_HEIGHT),
		table.WithWidth(DEFAULT_WIDTH),
	)
	style := table.DefaultStyles()
	style.Header = style.Header.Bold(true).
		BorderBottom(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
	style.Selected = style.Selected.
		Bold(true).
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("10"))

	coinTable.SetStyles(style)
	rateTable := table.New(
		table.WithHeight(DEFAULT_HEIGHT),
		table.WithWidth(DEFAULT_WIDTH),
	)
	rateTable.SetStyles(style)

	return SwapTable{
		coinTable: coinTable,
		rateTable: rateTable,
		state:     CoinTableState,
		api:       NewSwapAPI(client),
	}
}

func (t SwapTable) Init() tea.Cmd {
	t.rateTable.SetRows([]table.Row{})
	return tea.Batch(
		t.api.ListCoins(),
	)
}

func (t SwapTable) Update(msg tea.Msg) (SwapTable, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		var c tea.Cmd
		switch t.state {
		case CoinTableState:
			t.coinTable, c = t.coinTable.Update(msg)
		case RateTableState:
			t.rateTable, c = t.rateTable.Update(msg)
		}
		cmds = append(cmds, c)

	case tea.WindowSizeMsg: // not used yet
		t.coinTable.SetWidth(msg.Width)
		t.coinTable.SetHeight(msg.Height)
		t.rateTable.SetWidth(msg.Width)
		t.rateTable.SetHeight(msg.Height)

	case CoinFilterMsg:
		rows := t.FilterRows(msg.Query)
		t.coinTable.SetRows(rows)

	case CoinReqRespMsg:
		cmds = append(cmds, AddLog("table: coin req success"))
		t.coinList = msg.Resp
		rows := make([]table.Row, len(t.coinList))
		for i, coin := range t.coinList {
			rows[i] = table.Row{
				coin.GetName(),
				strings.ToLower(coin.GetTicker()),
				coin.GetNetwork(),
				fmt.Sprintf("%f", coin.GetMinimum()),
				fmt.Sprintf("%f", coin.GetMaximum()),
			}
		}
		t.coinTableRows = rows
		t.coinTable.SetRows(rows)

	case SwapRateReqMsg:
		cmds = append(cmds,
			AddLog("table: swap rate req"),
			t.api.SwapRate(msg),
		)

	case SwapRateRespMsg:
		cmds = append(cmds, AddLog("table: swap rate req success"))
		quotes := msg.Resp.Quotes.GetQuotes()
		ticker := func() string {
			var s string
			if t.payment {
				s = strings.ToUpper(msg.Resp.TickerFrom)
			} else {
				s = strings.ToUpper(msg.Resp.TickerTo)
			}
			return fmt.Sprintf("Amount (%s)", s)
		}()
		columns := []table.Column{
			{Title: "DEX Name", Width: 15},
			{Title: ticker, Width: 12},
			{Title: "USD", Width: 10},
			{Title: "Rate", Width: 8},
			{Title: "KYC", Width: 4},
			{Title: "ETA", Width: 8},
			{Title: "Spread", Width: 8},
		}
		sort.Slice(quotes, func(i, j int) bool {
			return quotes[i].GetAmountTo_USD() > quotes[j].GetAmountTo_USD()
		})
		rows := make([]table.Row, len(quotes))
		for i, quote := range quotes {
			rows[i] = table.Row{
				quote.GetProvider(),
				func() string {
					if t.payment {
						return quote.GetAmountFrom()
					}
					return quote.GetAmountTo()
				}(),
				func() string {
					if t.payment {
						return fmt.Sprintf("$%s", quote.GetAmountFrom_USD())
					}
					return fmt.Sprintf("$%s", quote.GetAmountTo_USD())
				}(),
				func() string {
					if quote.GetFixed() == "True" {
						return "Fixed"
					}
					return "Floating"
				}(),
				quote.GetKycrating(),
				fmt.Sprintf("%.f mins", quote.GetEta()),
				fmt.Sprintf("%s%%", quote.GetWaste()),
			}
		}
		t.state = RateTableState
		t.rateTable.SetColumns(columns)
		t.rateTable.SetRows(rows)
		t.rateTable.SetCursor(0)
		t.rateTable.Focus()

	case SwapTradeRespMsg:
		cmds = append(cmds, AddLog("table: SwapTradeResponse"))
	}

	return t, tea.Batch(cmds...)
}

func (t SwapTable) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8")).
		BorderBottom(false).
		Bold(true)
	switch t.state {
	case CoinTableState:
		return style.Render(t.coinTable.View())
	case RateTableState:
		return style.Render(t.rateTable.View())
	}
	return "table error"
}

func (t SwapTable) Blur() {
	switch t.state {
	case CoinTableState:
		t.coinTable.Blur()
	case RateTableState:
		t.rateTable.Blur()
	}
}

func (t SwapTable) Focus() tea.Cmd {
	switch t.state {
	case CoinTableState:
		t.coinTable.Focus()
	case RateTableState:
		t.rateTable.Focus()
	}
	return AddLog("table: focus %v", t.state)
}

func (t SwapTable) GetCoin() *pb.Coin {
	if t.coinTable.Cursor() >= len(t.coinList) {
		return nil
	}
	selected := t.coinTable.SelectedRow()
	for i, coin := range t.coinList {
		if coin.GetTicker() == selected[1] &&
			coin.GetNetwork() == selected[2] {
			return t.coinList[i]
		}
	}
	return nil
}

func (t SwapTable) GetExchange() string {
	if t.state != RateTableState {
		return ""
	}
	selected := t.rateTable.SelectedRow()
	return selected[0]
}

type CoinFilterMsg struct {
	Query string
}

func (t SwapTable) FilterRows(query string) []table.Row {
	AddLog("table: filter rows %s", query)
	filtered := []table.Row{}
	query = strings.ToLower(query)
	for _, row := range t.coinTableRows {
		if strings.Contains(strings.ToLower(row[0]), query) ||
			strings.Contains(strings.ToLower(row[1]), query) ||
			strings.Contains(strings.ToLower(row[2]), query) {
			filtered = append(filtered, row)
		}
	}
	return filtered
}

func (t SwapTable) Width() int {
	switch t.state {
	case CoinTableState:
		return t.coinTable.Width()
	case RateTableState:
		return t.rateTable.Width()
	}
	return 0
}

func (t SwapTable) Height() int {
	switch t.state {
	case CoinTableState:
		return t.coinTable.Height()
	case RateTableState:
		return t.rateTable.Height()
	}
	return 0
}

func (t SwapTable) SetPayment(v bool) SwapTable {
	t.payment = v
	return t
}
