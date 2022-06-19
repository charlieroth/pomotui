package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charlieroth/pomotui/state"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

func CreateView(m Model) string {
	view := GetTitle(m)

	switch m.State {
	case state.ChooseWorkingDuration, state.ChooseBreakDuration, state.ChooseLongBreakDuration, state.ChooseSessionCount:
		view += ChoicesView(m)
	case state.Working, state.Break, state.LongBreak:
		view += MainView(m)
	}

	view += HelpView(m)
	return view
}

func WorkingDurationTitle() string {
	return "Choose a working duration:\n"
}

func BreakDurationTitle() string {
	return "Choose a break duration:\n"
}

func LongBreakDurationTitle() string {
	return "Choose a long break duration:\n"
}
func SessionCountTitle() string {
	return "Choose number of sessions:\n"
}

func WorkingTitle() string {
	return "Work\n"
}

func BreakTitle() string {
	return "Break\n"
}

func LongBreakTitle() string {
	return "Long Break\n"
}

func GetTitle(m Model) string {
	switch m.State {
	case state.ChooseWorkingDuration:
		return WorkingDurationTitle()
	case state.ChooseBreakDuration:
		return BreakDurationTitle()
	case state.ChooseLongBreakDuration:
		return LongBreakDurationTitle()
	case state.ChooseSessionCount:
		return SessionCountTitle()
	case state.Working:
		return WorkingTitle()
	case state.Break:
		return BreakTitle()
	case state.LongBreak:
		return LongBreakTitle()
	}

	return "\n"
}

func RenderChoice(m Model, cursor, checked, choice string) string {
	switch m.State {
	case state.ChooseWorkingDuration, state.ChooseBreakDuration, state.ChooseLongBreakDuration:
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
	sessionCount, err := strconv.Atoi(m.SessionCount.selected)
	if err != nil {
		panic("failed convert session count from string to int")
	}

	var s strings.Builder
	for i := 1; i <= sessionCount; i++ {
		if m.CurrentWorkSession >= i {
			activeDot := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
			s.WriteString(" " + activeDot)
		} else {
			inActiveDot := lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
			s.WriteString(" " + inActiveDot)
		}
	}
	view += fmt.Sprintf("\n%s", s.String())
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
