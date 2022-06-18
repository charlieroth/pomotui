package model

import (
	"fmt"

	"github.com/charlieroth/pomotui/state"
	"github.com/charmbracelet/bubbles/key"
)

func CreateView(m Model) string {
	view := ""
	view += GetPrompt(m)

	switch m.State {
	case state.ChooseWorkingDuration, state.ChooseBreakDuration, state.ChooseSessionCount:
		view += ChoicesView(m)
	case state.Working, state.Break:
		view += MainView(m)
	}

	view += HelpView(m)
	return view
}

func WorkingDurationPrompt() string {
	return "Choose a working duration:\n"
}

func ChooseBreakDurationPrompt() string {
	return "Choose a break duration:\n"
}

func ChooseSessionCountPrompt() string {
	return "Choose number of sessions:\n"
}

func WorkingPrompt() string {
	return "Working\n"
}

func BreakPrompt() string {
	return "Break :)\n"
}

func GetPrompt(m Model) string {
	switch m.State {
	case state.ChooseWorkingDuration:
		return WorkingDurationPrompt()
	case state.ChooseBreakDuration:
		return ChooseBreakDurationPrompt()
	case state.ChooseSessionCount:
		return ChooseSessionCountPrompt()
	case state.Working:
		return WorkingPrompt()
	case state.Break:
		return BreakPrompt()
	}

	return "\n"
}

func RenderChoice(m Model, cursor, checked, choice string) string {
	switch m.State {
	case state.ChooseWorkingDuration:
		return fmt.Sprintf("%s [%s] %s mins\n", cursor, checked, choice)
	case state.ChooseBreakDuration:
		return fmt.Sprintf("%s [%s] %s mins\n", cursor, checked, choice)
	case state.ChooseSessionCount:
		return fmt.Sprintf("%s [%s] %s sessions\n", cursor, checked, choice)
	}

	return ""
}

func ChoicesView(m Model) string {
	view := ""

	currentCursor := m.CurrentCursor()
	currentSelectedChoice := m.CurrentSelectedChoice()
	currentChoices := m.CurrentChoices()

	for i, choice := range currentChoices {
		cursor := " "
		if currentCursor == i {
			cursor = ">"
		}

		checked := " "
		if currentSelectedChoice == choice {
			checked = "x"
		}

		view += RenderChoice(m, cursor, checked, choice)
	}

	return view
}

func MainView(m Model) string {
	view := ""
	view += m.Timer.View()
    currentSession := m.CurrentWorkSession
    totalWorkingSessions := m.SessionCount.selected
	view += fmt.Sprintf("\n%d of %s", currentSession, totalWorkingSessions)
	return view
}

func HelpView(m Model) string {
	return "\n" + m.Help.ShortHelpView([]key.Binding{
		m.KeyMap.Up,
		m.KeyMap.Down,
		m.KeyMap.Enter,
		m.KeyMap.Confirm,
		m.KeyMap.Start,
		m.KeyMap.Stop,
		m.KeyMap.Reset,
		m.KeyMap.Quit,
	})
}
