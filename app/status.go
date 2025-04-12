package app

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	pb "github.com/lfaoro/swap/gen/go/swap/v1"
	"github.com/skip2/go-qrcode"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (m *SwapUI) formatStatus(s string) string {
	var out string
	switch s {
	case "waiting":
		out += fmt.Sprintf("[waiting for %s block]", m.cd.From.GetName())
	case "confirming":
		out += fmt.Sprintf("[confirming %s block]", m.cd.From.GetName())
	case "sending":
		out += fmt.Sprintf("[sending to %s address]", m.cd.To.GetName())
	case "finished":
		out += fmt.Sprintf("[transfer successful to %s address]\nThank you for using Swap!", m.cd.To.GetName())
	case "paid_partially":
		out += fmt.Sprintf("[partially paid]")
	case "failed":
		out += fmt.Sprintf("[failed]")
	case "expired":
		out += fmt.Sprintf("[expired]")
	case "halted":
		out += fmt.Sprintf("[halted]")
	case "refunded":
		out += fmt.Sprintf("[refunded]")
	}
	return out
}

func (m *SwapUI) setStatusContent(msg SwapStatusRespMsg) string {
	data := msg.SwapStatusResponse
	type Status struct {
		Status string
		Date   string

		TradeID      string
		IdProvider   string
		Provider     string
		ProviderLogo string
		OriginalEta  int

		Fixed  bool
		Market string

		TickerFrom  string
		TickerTo    string
		NetworkFrom string
		NetworkTo   string
		CoinFrom    string
		CoinTo      string
		AmountFrom  float64
		AmountTo    float64

		AddressUser     string
		AddressProvider string

		SupportURL string
		TimeLeft   time.Duration
		Progress   string

		QRData string
		QRCode string

		Loading string
	}
	// NOTE: this is the time format used by the API
	// 2025-03-19T11:16:02.958665Z
	const timeFormat = "2006-01-02T15:04:05.000000Z"
	status := Status{
		Date:            data.Date.AsTime().Format(timeFormat),
		TradeID:         data.GetTradeId(),
		IdProvider:      data.GetIdProvider(),
		Provider:        data.GetProvider(),
		Fixed:           data.GetFixed(),
		NetworkFrom:     data.GetNetworkFrom(),
		NetworkTo:       data.GetNetworkTo(),
		CoinFrom:        data.GetCoinFrom(),
		CoinTo:          data.GetCoinTo(),
		AmountFrom:      data.GetAmountFrom(),
		AmountTo:        data.GetAmountTo(),
		AddressUser:     data.GetAddressUser(),
		AddressProvider: data.GetAddressProvider(),
	}
	status.Status = m.formatStatus(data.Status)
	status.Loading = m.sp.View()

	status.QRData = fmt.Sprintf("%s:%s?amount=%f", strings.ToLower(data.CoinFrom), strings.TrimSpace(data.AddressProvider), data.AmountFrom)
	qr, err := qrcode.New(status.QRData, qrcode.Low)
	if err != nil {
		return fmt.Sprintf("error %v", err)
	}
	qr.DisableBorder = true
	status.QRCode = qr.ToSmallString(false)

	if data.Details != nil {
		status.SupportURL = data.Details.Support.TxUrl
		status.OriginalEta = int(data.Details.OriginalEta)
		status.TickerFrom = strings.ToUpper(data.TickerFrom)
		status.TickerTo = strings.ToUpper(data.TickerTo)

		expiresAt, err := time.Parse(timeFormat, data.Details.ExpiresAt.AsTime().Format(timeFormat))
		if err != nil {
			return fmt.Sprintf("error %v", err)
		}
		status.TimeLeft = expiresAt.Sub(time.Now()).Round(time.Minute)

		creationTime, err := time.Parse(timeFormat, data.Date.AsTime().Format(timeFormat))
		if err != nil {
			return fmt.Sprintf("error parsing creation time: %v", err)
		}
		totalDuration := expiresAt.Sub(creationTime).Seconds()
		timeLeft := expiresAt.Sub(time.Now()).Seconds()
		timeLeftPrc := timeLeft / totalDuration
		if timeLeftPrc < 0 {
			timeLeftPrc = 0
		} else if timeLeftPrc > 1 {
			timeLeftPrc = 1
		}
		status.Progress = m.pg.ViewAs(timeLeftPrc)
	}

	m.sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	m.SetSpinning(true)

	templ := `
	Status 
	→{{.Status}} {{.Loading}}

	{{.CoinFrom}} ({{.NetworkFrom}}) [{{if .Fixed}}FIXED{{else}}FLOAT{{end}}] {{.CoinTo}} ({{.NetworkTo}})
	{{.AmountFrom}} {{.TickerFrom}} -> {{.AmountTo}} {{.TickerTo}}

	Send {{.AmountFrom}} {{.TickerFrom}} to
	[{{.AddressProvider}}]

	Receive {{.AmountTo}} {{.TickerTo}} in ~{{.OriginalEta}} minutes on
	[{{.AddressUser}}]

	Expires in {{.TimeLeft}} [{{.Progress}}]

	Support ({{.SupportURL}})
	TradeID {{.IdProvider}}
	Provider {{.Provider}}
	Chat (https://t.me/swapcli)

Can use this QR code to send {{.AmountFrom}} {{.TickerFrom}} to
{{.AddressProvider}}

{{.QRCode}}
	`

	t := template.Must(template.New("status").Parse(templ))
	var out strings.Builder
	err = t.Execute(&out, status)
	if err != nil {
		return fmt.Sprintf("error %v", err)
	}
	return out.String()
}

func (m *SwapUI) debugStatusData() SwapStatusRespMsg {
	m.cd.From = &pb.Coin{
		Name:    "Bitcoin",
		Network: "Mainnet",
		Ticker:  "btc",
	}
	m.cd.To = &pb.Coin{
		Name:    "Monero",
		Network: "Mainnet",
		Ticker:  "xmr",
	}

	return SwapStatusRespMsg{
		SwapStatusResponse: &pb.SwapStatusResponse{
			Status:              "confirming",
			TradeId:             "Z0HC0YNfag",
			Date:                timestamppb.New(time.Now()),
			Type:                "api",
			TickerFrom:          "btc",
			TickerTo:            "xmr",
			CoinFrom:            "Bitcoin",
			CoinTo:              "Monero",
			NetworkFrom:         "Mainnet",
			NetworkTo:           "Mainnet",
			AmountFrom:          0.00255308,
			AmountTo:            1.0,
			Provider:            "FixedFloat",
			Payment:             true,
			Fixed:               true,
			AddressProvider:     "bc1q80pg5gn5ynev5eveauq445fey5cr9l8cn4dhua",
			AddressProviderMemo: "",
			AddressUser:         "89XCyahmZiQgcVwjrSZTcJepPqCxZgMqwbABvzPKVpzC7gi8URDme8H6UThpCqX69y5i1aA81AKq57Wynjovy7g4K9MeY5c",
			AddressUserMemo:     "",
			RefundAddress:       "",
			RefundAddressMemo:   "",
			Password:            "upsaofbiBg312D5CPFvPxGMkC7nHuRx1WJyCAEeY",
			IdProvider:          "44AS61",
			Details: &pb.Details{
				Hashout: "",
				Support: &pb.Support{
					TxUrl:      "https://ff.io/order/44AS61",
					SupportUrl: "https://ff.io/support",
					Tos:        "https://ff.io/tos",
				},
				Webhook:            "",
				ExpiresAt:          timestamppb.New(time.Now().Add(time.Minute * 10)),
				AmountBtc:          0,
				OriginalEta:        4.0,
				ProviderLogo:       "",
				MarketrateCreation: 397.09723784761144,
			},
		},
	}
}
