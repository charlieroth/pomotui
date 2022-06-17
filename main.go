package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ChooseWorkingDuration = "ChooseWorkingDuration"
	ChooseBreakDuration   = "ChooseBreakDuration"
	ChooseSessionCount    = "ChooseSessionCount"
	Working               = "Working"
	Break                 = "Break"
)

type WorkingDuration struct {
	choices  []string
	cursor   int
	selected string
}

type BreakDuration struct {
	choices  []string
	cursor   int
	selected string
}

type SessionCount struct {
	choices  []string
	cursor   int
	selected string
}

type State struct {
	Step                           string
	CurrentWorkSession             int
	WorkingDurationTimer           timer.Model
	HasStartedWorkingDurationTimer bool
	BreakDurationTimer             timer.Model
	HasStartedBreakDurationTimer   bool
	WorkingDuration                WorkingDuration
	BreakDuration                  BreakDuration
	SessionCount                   SessionCount
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
		Step:               ChooseWorkingDuration,
		CurrentWorkSession: 0,
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
		}

		selectedTime, err := strconv.Atoi(s.WorkingDuration.selected)
		if err != nil {
			panic("Failed to convert working duration time to int")
		}

		amountOfTime := time.Duration(selectedTime) * time.Minute
		s.WorkingDurationTimer = timer.NewWithInterval(amountOfTime, time.Second)
		s.Step = ChooseBreakDuration
	case ChooseBreakDuration:
		if !s.HasSelectedBreakDuration() {
			return s, nil
		}

		selectedTime, err := strconv.Atoi(s.BreakDuration.selected)
		if err != nil {
			panic("Failed to convert break duration time to int")
		}

		amountOfTime := time.Duration(selectedTime) * time.Second
		s.BreakDurationTimer = timer.New(amountOfTime)
		s.Step = ChooseSessionCount
	case ChooseSessionCount:
		if !s.HasSelectedSessionCount() {
			return s, nil
		}

		sessionCount, err := strconv.Atoi(s.SessionCount.selected)
		if err != nil {
			panic("Failed to convert session count to int")
		}

		s.CurrentWorkSession = sessionCount
		s.Step = Working
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

func HandleWorkingDurationTimerTickMsg(s State, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	s.WorkingDurationTimer, cmd = s.WorkingDurationTimer.Update(msg)
	return s, cmd
}

func HandleWorkingDurationTimerStartStopMsg(s State, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	s.WorkingDurationTimer, cmd = s.WorkingDurationTimer.Update(msg)
	return s, cmd
}

func (s State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		return HandleWorkingDurationTimerTickMsg(s, msg)
	case timer.StartStopMsg:
		return HandleWorkingDurationTimerStartStopMsg(s, msg)
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
		case "s":
			return s, s.WorkingDurationTimer.Toggle()
		case "i":
            s.HasStartedWorkingDurationTimer = true
			return s, s.WorkingDurationTimer.Init()
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

func WorkingPrompt() string {
	return "Working\n"
}

func BreakPrompt() string {
	return "Break :)\n"
}

func GetPrompt(s State) string {
	switch s.Step {
	case ChooseWorkingDuration:
		return WorkingDurationPrompt()
	case ChooseBreakDuration:
		return ChooseBreakDurationPrompt()
	case ChooseSessionCount:
		return ChooseSessionCountPrompt()
	case Working:
		return WorkingPrompt()
	case Break:
		return BreakPrompt()
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

func WorkingView(s State) string {
	view := ""
	view += s.WorkingDurationTimer.View()
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
	case Working:
		if s.HasStartedWorkingDurationTimer && s.WorkingDurationTimer.Running() {
			view += ", s = stop\n"
			return view
		}

		if s.WorkingDurationTimer.Timedout() {
			// TODO(charlieroth): show "start break" message
			view += ", timedout\n"
			return view
		}

		if s.HasStartedWorkingDurationTimer && !s.WorkingDurationTimer.Running() {
            view += ", s = start\n"
        } else {
            view += ", i = init\n"
        }
		return view
	}

	view += "\n"
	return view
}

func (s State) View() string {
	view := ""
	view += GetPrompt(s)

	switch s.Step {
	case ChooseWorkingDuration, ChooseBreakDuration, ChooseSessionCount:
		view += ChoicesView(s)
	case Working:
		view += WorkingView(s)
	}

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
