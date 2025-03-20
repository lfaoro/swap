package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Help struct {
	color lipgloss.Color
}

func NewHelp() Help {
	return Help{
		color: lipgloss.Color("8"),
	}
}

func (h Help) Init() tea.Cmd {
	return nil
}

func (h Help) Update(msg tea.Msg) (Help, tea.Cmd) {
	return h, nil
}

func (h Help) View() string {
	text := "↑/↓ navigate • enter select • / search • esc cancel • tab swap/pay • q quit"
	return lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(h.color).
		Render(text)
}

func (h Help) AddressHelp() string {
	text := " ↑ reuse address • ctrl+d delete address • ctrl+s save address"
	return lipgloss.NewStyle().
		Foreground(h.color).
		Render(text)
}
