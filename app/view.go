package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *SwapUI) View() string {
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
			m.trxView.Style = lipgloss.Style{}
			titleStyle := func() string {
				b := lipgloss.RoundedBorder()
				b.Right = "├"
				style := lipgloss.NewStyle().
					Border(b).
					Padding(0, 1)
				title := style.Render("Transaction")
				line := strings.Repeat("─", max(0, m.trxView.Width-lipgloss.Width(title)))
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
				info := style.Render(fmt.Sprintf("%2.f%%", m.trxView.ScrollPercent()*100))
				line := strings.Repeat("─", max(0, m.trxView.Width-lipgloss.Width(info)))
				return lipgloss.JoinHorizontal(
					lipgloss.Center,
					line,
					info,
				)
			}()

			return titleStyle + "\n" + m.trxView.View() + "\n" + infoStyle
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
