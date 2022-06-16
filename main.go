package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	ChooseWorkingDuration = "ChooseWorkingDuration"
	ChooseBreakDuration   = "ChooseBreakDuration"
	ChooseSessionCount    = "ChooseSessionCount"
	Idle                  = "Idle"
	Working               = "Working"
)

// working duration: 15, 20, 25, 30, 45, 50, 60, 90
// break duration: 5, 10, 15, 20
// session count: 4, 5, 6, 7

type WorkingDuration struct {
	choices  []string // options in working duration
	cursor   int      // which working duration option cursor is point at
	selected string   // which working duration was selected
}

type BreakDuration struct {
	choices  []string // options in working duration
	cursor   int      // which working duration option cursor is point at
	selected string   // which working duration was selected
}

type SessionCount struct {
	choices  []string // options in working duration
	cursor   int      // which working duration option cursor is point at
	selected string   // which working duration was selected
}

type State struct {
	Step            string
	WorkingDuration WorkingDuration
	BreakDuration   BreakDuration
	SessionCount    SessionCount
}

func (s State) HasSelectedWorkingDuration() bool {
	return s.WorkingDuration.selected != ""
}

func (s State) HasSelectedBreakDuration() bool {
	return s.BreakDuration.selected != ""
}

func (s State) HasSelectedSessionCount() bool {
	return s.SessionCount.selected != ""
}

func (s State) CurrentCursor() int {
    switch s.Step {
    case ChooseWorkingDuration:
        return s.WorkingDuration.cursor
    case ChooseBreakDuration:
        return s.BreakDuration.cursor
    case ChooseSessionCount:
        return s.SessionCount.cursor
    }

    return 0
}

func (s State) CurrentSelectedChoice() string {
    switch s.Step {
    case ChooseWorkingDuration:
        return s.WorkingDuration.selected
    case ChooseBreakDuration:
        return s.BreakDuration.selected
    case ChooseSessionCount:
        return s.SessionCount.selected
    }

    return ""
}

func (s State) CurrentChoices() []string {
    switch s.Step {
    case ChooseWorkingDuration:
        return s.WorkingDuration.choices
    case ChooseBreakDuration:
        return s.BreakDuration.choices
    case ChooseSessionCount:
        return s.SessionCount.choices
    }

    return []string{}
}

func InitState() State {
	return State{
		Step: ChooseWorkingDuration,
		WorkingDuration: WorkingDuration{
			choices:  []string{"15", "20", "25", "30", "45", "50", "60", "90"},
			cursor:   0,
			selected: "",
		},
		BreakDuration: BreakDuration{
			choices:  []string{"5", "10", "15", "20"},
			cursor:   0,
			selected: "",
		},
		SessionCount: SessionCount{
			choices:  []string{"4", "5", "6", "7"},
			cursor:   0,
			selected: "",
		},
	}
}

func (s State) Init() tea.Cmd {
	return nil
}

func HandleQuit(s State) (tea.Model, tea.Cmd) {
	return s, tea.Quit
}

func HandleUp(s State) (tea.Model, tea.Cmd) {
    switch s.Step {
	case ChooseWorkingDuration:
        if s.WorkingDuration.cursor > 0 {
            s.WorkingDuration.cursor--
        }
        return s, nil
	case ChooseBreakDuration:
        if s.BreakDuration.cursor > 0 {
            s.BreakDuration.cursor--
        }
        return s, nil
	case ChooseSessionCount:
        if s.SessionCount.cursor > 0 {
            s.SessionCount.cursor--
        }
        return s, nil
	}

	return s, nil
}

func HandleDown(s State) (tea.Model, tea.Cmd) {
    switch s.Step {
	case ChooseWorkingDuration:
        if s.WorkingDuration.cursor < len(s.WorkingDuration.choices)-1 {
            s.WorkingDuration.cursor++
        }
        return s, nil
	case ChooseBreakDuration:
        if s.BreakDuration.cursor < len(s.BreakDuration.choices)-1 {
            s.BreakDuration.cursor++
        }
        return s, nil
	case ChooseSessionCount:
        if s.SessionCount.cursor < len(s.SessionCount.choices)-1 {
            s.SessionCount.cursor++
        }
        return s, nil
	}

	return s, nil
}

func HandleConfirm(s State) (tea.Model, tea.Cmd) {
	switch s.Step {
	case ChooseWorkingDuration:
		if !s.HasSelectedWorkingDuration() {
			return s, nil
		} else {
			s.Step = ChooseBreakDuration
		}
	case ChooseBreakDuration:
		if !s.HasSelectedBreakDuration() {
			return s, nil
		} else {
			s.Step = ChooseSessionCount
		}
	case ChooseSessionCount:
		if !s.HasSelectedSessionCount() {
			return s, nil
		} else {
			s.Step = Idle
		}
	}
	return s, nil
}

func HandleEnter(s State) (tea.Model, tea.Cmd) {
    switch s.Step {
	case ChooseWorkingDuration:
        if s.WorkingDuration.selected == "" {
            s.WorkingDuration.selected = s.WorkingDuration.choices[s.WorkingDuration.cursor]
        } else {
            if s.WorkingDuration.choices[s.WorkingDuration.cursor] != s.WorkingDuration.selected {
                s.WorkingDuration.selected = s.WorkingDuration.choices[s.WorkingDuration.cursor]
            } else {
                s.WorkingDuration.selected = ""
            }
        }
        return s, nil
	case ChooseBreakDuration:
        if s.BreakDuration.selected == "" {
            s.BreakDuration.selected = s.BreakDuration.choices[s.BreakDuration.cursor]
        } else {
            if s.BreakDuration.choices[s.BreakDuration.cursor] != s.BreakDuration.selected {
                s.BreakDuration.selected = s.BreakDuration.choices[s.BreakDuration.cursor]
            } else {
                s.BreakDuration.selected = ""
            }
        }
        return s, nil
	case ChooseSessionCount:
        if s.SessionCount.selected == "" {
            s.SessionCount.selected = s.SessionCount.choices[s.SessionCount.cursor]
        } else {
            if s.SessionCount.choices[s.SessionCount.cursor] != s.SessionCount.selected {
                s.SessionCount.selected = s.SessionCount.choices[s.SessionCount.cursor]
            } else {
                s.SessionCount.selected = ""
            }
        }
        return s, nil
	}

	return s, nil
}

func (s State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return HandleQuit(s)
		case "up", "k":
			return HandleUp(s)
		case "down", "j":
			return HandleDown(s)
		case "enter", " ":
			return HandleEnter(s)
		case "c":
			return HandleConfirm(s)
		}
	}

	return s, nil
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

func IdlePrompt() string {
	return "Get to work!\n"
}

func GetPrompt(s State) string {
	switch s.Step {
	case ChooseWorkingDuration:
		return WorkingDurationPrompt()
	case ChooseBreakDuration:
		return ChooseBreakDurationPrompt()
	case ChooseSessionCount:
		return ChooseSessionCountPrompt()
	case Idle:
		return IdlePrompt()
	}

	return "\n"
}

func RenderChoice(s State, cursor, checked, choice string) string {
	switch s.Step {
	case ChooseWorkingDuration:
        return fmt.Sprintf("%s [%s] %s mins\n", cursor, checked, choice)
	case ChooseBreakDuration:
        return fmt.Sprintf("%s [%s] %s mins\n", cursor, checked, choice)
	case ChooseSessionCount:
        return fmt.Sprintf("%s [%s] %s sessions\n", cursor, checked, choice)
	}

    return ""
}

func ChoicesView(s State) string {
	view := ""

    currentCursor := s.CurrentCursor()
    currentSelectedChoice := s.CurrentSelectedChoice()
    currentChoices := s.CurrentChoices()

	for i, choice := range currentChoices {
		cursor := " "
		if currentCursor == i {
			cursor = ">"
		}

		checked := " "
		if currentSelectedChoice == choice {
			checked = "x"
		}

		view += RenderChoice(s, cursor, checked, choice)
	}

	return view
}

func FooterView(s State) string {
	view := "\nq = quit"

	switch s.Step {
	case ChooseWorkingDuration:
		if s.HasSelectedWorkingDuration() {
			view += ", c = confirm\n"
			return view
		}
	case ChooseBreakDuration:
		if s.HasSelectedBreakDuration() {
			view += ", c = confirm\n"
			return view
		}
	case ChooseSessionCount:
		if s.HasSelectedSessionCount() {
			view += ", c = confirm\n"
			return view
		}
	case Idle:
		if s.HasSelectedSessionCount() {
			view += ", s = start, s = stop\n"
			return view
		}
	}

	view += "\n"
	return view
}

func (s State) View() string {
	view := ""
	view += GetPrompt(s)
	view += ChoicesView(s)
	view += FooterView(s)
	return view
}

func main() {
	p := tea.NewProgram(InitState())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
