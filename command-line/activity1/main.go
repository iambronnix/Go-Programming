package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)
type Units struct{
	abstract int
    linear   int
    iap      int
   statistics int	
}
type Individual struct {
	unit  []Units
	name []string
	department string
}

type Model struct {
	student []Individual
	Cursor   int
}

func initialModel() Model {
	return Model{
		choices:  []string{"File input", "Type in input"},
		selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
}

func (m Model) View() string {
	s := "Choose an option:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
