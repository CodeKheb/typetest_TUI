package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

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

	menuStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#666666")).
			Padding(10, 10).
			Bold(true)

	wpmStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#666666")).
			Padding(0, 1)
	
	finishedWPM = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#666666")).
			Padding(0, 1)
	
	dimmedStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#444444")).
			Padding(0, 1)
)

var hintText string = "Esc: QUIT | Enter: RESET | Tab: MENU"
func (m Model) menuText() string {
	best := bestScore(m.leaderboard)
	return fmt.Sprintf(
		"MENU\n"+
		"Typing Test Written in Go\n\n"+
		"Best WPM: %.1f\n"+
		"Last Run: %.1f WPM\n\n"+
		"1 - 9: Change Word Count\n"+
		"Tab: Type\n"+
		"Esc: Quit",
		best, m.WPM,
	)
}

