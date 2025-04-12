package app

import tea "github.com/charmbracelet/bubbletea"

type Input struct {
	amount  string
	address string
}

func NewInput() *Input {
	return &Input{
		amount:  "",
		address: "",
	}
}

func (i *Input) Init() tea.Cmd {
	return nil
}

func (i *Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return i, nil
}

func (i *Input) View() string {
	return i.amount + " " + i.address
}
