package main

import (
	"fmt"
	"os"
	"time"

	// use tea bubbletea and lipgloss for UI
	tea "github.com/charmbracelet/bubbletea"
)


type Screen int

const (
	TypingScreen Screen = iota
	MenuScreen
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
	timeStart time.Time
	started bool
	finished bool
	WPM float64

	screen Screen
	leaderboard []LeaderboardEntry
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
		return m, nil

	case tea.KeyMsg:
		switch m.screen {
		case MenuScreen:
			return m.updateMenu(message)
		case TypingScreen:
			return m.updateTyping(message)
		}
	}
	return m, nil
}

func (m Model) currentWPM() float64 {
	if !m.started {
		return 0
	}

	elapsed := time.Since(m.timeStart).Minutes()
	if elapsed <= 0 {
		return 0
	}	
	
	// counts "words" per 5 characters divided by elapsed time to calculate the WPM
	return float64(len(m.typed)) / 5.0 / elapsed	
}


// Lipgloss UI in tea View
func (m Model) View() string {
	switch m.screen {
	case MenuScreen:
		return m.menuView()
	default:
		return m.typingView()
	}
}

// Main
func main() {
	leaderboard, _ := loadLeaderboard()

	m := Model{
		target: randomizedWords(),
		leaderboard: leaderboard,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
