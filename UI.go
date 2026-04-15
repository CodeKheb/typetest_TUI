package main

import "github.com/charmbracelet/lipgloss"

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
var menuText string = "MENU\n" + 
"Typing Test Written in Go" + 
"\n1: 10 words\n2: 20 words\n3: 30 words\n4: 40 words\n5: 50 words\n6: 60 words\n7: 70 words\n8: 80 words\n9: 90 words"

