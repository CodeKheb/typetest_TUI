package main

import (
	"fmt"
	"os"

	// use tea bubbletea
	tea "github.com/charmbracelet/bubbletea"
)

/* Model struct
TO DO: Make Target Random, add UI live for target
make timer and show wrong types
add a accuracy calculator in the future as well
*/
type Model struct {
	typed string
	target string
}

// Initialiazes Model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update Logic switch statement for message tea.Msg and case ctrl+c or esc quit the program
// backspace logic, gets the length of the current typed and minus 1
// default just takes user input and puts it in message.String
func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.KeyMsg:

		switch message.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "backspace":
			if len(m.typed) > 0 {
				m.typed = m.typed[:len(m.typed)-1]
			}
		default: 
			m.typed += message.String()
		}
	}
	return m, nil
}

// View Logic, just a string parser for now, maybe use lipgloss here in the future for UI
func (m Model) View() string {
	return fmt.Sprintf(
		"%s\n%s\n\nPress Esc to Quit",
		m.target,
		m.typed,
	)
}

// Main
func main() {
	m := Model{
		target: "the quick brown fox jumps over the lazy dog",
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}


