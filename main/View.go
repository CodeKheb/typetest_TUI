package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) typingView() string {
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
				if char == " " {
					STRING_BUILDER.WriteString(wrongTyped.Render(string(m.typed[i])))
				} else {
					STRING_BUILDER.WriteString(wrongTyped.Render(char))
				}
			}
		} else if i == len(m.typed) {
			STRING_BUILDER.WriteString(cursor.Render(char))
		} else {
			STRING_BUILDER.WriteString(waiting.Render(char))
		}
	}

	// box text wrapper
	box := boxStyle.Width(m.width - 4).Render(STRING_BUILDER.String())

	// display WPM
	var displayWPM string
	if m.finished {
		displayWPM = finishedWPM.Render(fmt.Sprintf("FINAL WPM: %.0f", m.WPM))
	} else if m.started {
		displayWPM = wpmStyle.Render(fmt.Sprintf("CURRENT WPM: %.0f", m.currentWPM()))
	} else {
		displayWPM = dimmedStyle.Render("start typing... ")
	}

	// center the box
	pad := max(m.height/2 - 3, 0)
	verticalPad := strings.Repeat("\n", pad)

	// just the hint 
	hint := dimmedStyle.Render(hintText)

	return fmt.Sprintf(
		"%s%s\n\n%s\n%s",
		verticalPad,
		box,
		displayWPM,
		hint,
	)

}

func (m Model) updateTyping(message tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch message.String() {

	case "ctrl+c", "esc":
		return m, tea.Quit

	case "tab":
		m.screen = MenuScreen
		return m, nil

	case "backspace":
		if len(m.typed) > 0 {
			m.typed = m.typed[:len(m.typed)-1]
		}

	case "enter":
		return m.reset(), nil

	default:
		if !m.started {
			m.started = true
			m.timeStart = time.Now()
		}

		m.typed += message.String()

		if m.typed == m.target {
			elapsed := time.Since(m.timeStart).Minutes()
			words := float64(len(strings.Fields(m.target)))
			m.WPM = words / elapsed
			m.finished = true
		}
	}

	return m, nil
}

func (m Model) menuView() string {
	menu := menuStyle.Width(m.width - 4).Height(m.height - 4).Render(menuText)

	pad := max(m.height/2 - 3, 0)
	verticalPad := strings.Repeat("\n", pad)

	return fmt.Sprintf("%s%s", verticalPad, menu)
}


func (m Model) updateMenu(message tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch message.String() {
	case "tab":
		m.screen = TypingScreen
		return m, nil
	case "1":
		wordAmount = 10
		return m.reset(), nil
	case "2":
		wordAmount = 20
		return m.reset(), nil
	case "3":
		wordAmount = 30
		return m.reset(), nil
	case "4":
		wordAmount = 40
		return m.reset(), nil
	case "5":
		wordAmount = 50
		return m.reset(), nil
	case "6":
		wordAmount = 60
		return m.reset(), nil
	case "7":
		wordAmount = 70
		return m.reset(), nil
	case "8":
		wordAmount = 80
		return m.reset(), nil
	case "9":
		wordAmount = 90
		return m.reset(), nil
	case "ctrl+c", "esc":
		return m, tea.Quit
	}
	return m, nil
}
