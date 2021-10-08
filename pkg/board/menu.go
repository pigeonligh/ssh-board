package board

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
	username string
	choices  []MenuItem
	cursor   int
	ok       bool
}

func NewMenu(username string, boards []MenuItem) *MenuModel {
	return &MenuModel{
		username: username,
		choices:  boards,
		cursor:   0,
		ok:       false,
	}
}

func (m *MenuModel) Init() tea.Cmd {
	return nil
}

func (m *MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.ok = false
			return m, tea.Quit

		case "enter", " ":
			m.ok = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		}
	}

	return m, nil
}

func (m *MenuModel) View() string {
	s := fmt.Sprintf("Hello, %s!\n\nWhat do you want to play?\n\n", m.username)

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice.Name())
	}

	s += "\nPress q to quit.\n"
	return s
}

func (m *MenuModel) IsOK() bool {
	return m.ok
}

func (m *MenuModel) Run() {
	m.choices[m.cursor].RunOrDie(m.username)
}
