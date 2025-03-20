package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ChoicePrompt struct {
	active   bool
	choices  []string
	selected string
	cursor   int
}

// ensure interface compliance.

func NewChoicePrompt(choices []string) ChoicePrompt {
	return ChoicePrompt{
		choices: choices,
		cursor:  0,
	}
}

func (c ChoicePrompt) Toggle() (ChoicePrompt, tea.Cmd) {
	c.active = !c.active
	return c, AddLog("choice: active %v", c.active)
}
func (c ChoicePrompt) Active() bool {
	return c.active
}

func (c ChoicePrompt) Selected() string {
	return c.selected
}

// Init satisfied the tea.Model interface.
func (c ChoicePrompt) Init() tea.Cmd {
	return nil
}

func (c ChoicePrompt) Update(msg tea.Msg) (ChoicePrompt, tea.Cmd) {
	if !c.active {
		return c, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			c.active = false
			return c, AddLog("choice: cancelled")
		case "enter":
			c.selected = c.choices[c.cursor]
			c.active = false
			return c, AddLog("choice: selected %v", c.selected)
		case "left", "h":
			if c.cursor > 0 {
				c.cursor--
			}
			return c, nil
		case "right", "l":
			if c.cursor < len(c.choices)-1 {
				c.cursor++
			}
			return c, nil
		}
	}
	return c, nil
}
func (c ChoicePrompt) View() string {
	if !c.active {
		return ""
	}
	var out string
	for i, choice := range c.choices {
		var sel = ""
		if c.cursor == i {
			sel = ">"
		}
		out += fmt.Sprintf("%s%s ", sel, choice)
	}
	return out
}
