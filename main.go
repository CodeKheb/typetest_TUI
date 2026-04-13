package main

import (
	"fmt"
	"os"
	"strings"

	// use tea bubbletea and lipgloss for UI
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// COLORS FOR UI
var (
	correctTyped = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	wrongTyped = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	waiting = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	cursor = lipgloss.NewStyle().Underline(true)

	// box around the text
	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#666666")).
			Padding(0, 1)
)

/* Model struct
TO DO: Make Target Random, add UI live for target
make timer and show wrong types
add a accuracy calculator in the future as well
*/
type Model struct {
	typed string
	target string
	width int
	height int
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
	case tea.WindowSizeMsg:
		m.width = message.Width
		m.height = message.Height
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


// Lipgloss UI in tea View
func (m Model) View() string {
	// string builder
	var STRING_BUILDER strings.Builder

	// declare char as string for the target
	/* everything in the range of target(confusing name omg)
	if target is less than the typed(cannot go over)
	then if string is equal to char: render correctTyped
	else: render wrongTyped
	else if target is == length of typed: render cursor
	else: render waiting
	*/ 
	for i, target := range m.target {
		char := string(target)

		if i < len(m.typed) {
			if string(m.typed[i]) == char {
				STRING_BUILDER.WriteString(correctTyped.Render(char))
			} else {
				STRING_BUILDER.WriteString(wrongTyped.Render(char))
			}
		} else if i == len(m.typed) {
			STRING_BUILDER.WriteString(cursor.Render(char))
		} else {
			STRING_BUILDER.WriteString(waiting.Render(char))
		}
	}

	// box text wrapper
	box := boxStyle.Width(m.width - 4).Render(STRING_BUILDER.String())

	// center the box
	pad := max(m.height/2 - 3, 0)
	verticalPad := strings.Repeat("\n", pad)
	

	return fmt.Sprintf(
		"%s%s\n\n%s",
		verticalPad,
		box,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Render("ESC to quit"),
	)
}


// Main
func main() {
	m := Model{
		target: "the quick brown fox jumps over the lazy dog",
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
