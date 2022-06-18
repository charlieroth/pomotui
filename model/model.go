package model

import (
	"github.com/charlieroth/pomotui/state"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type WorkingDuration struct {
	choices  []string
	cursor   int
	selected string
}

func NewWorkingDuration() WorkingDuration {
	return WorkingDuration{
		choices:  []string{"15", "20", "25", "30", "45", "50", "60", "90"},
		cursor:   0,
		selected: "",
	}
}

type BreakDuration struct {
	choices  []string
	cursor   int
	selected string
}

func NewBreakDuration() BreakDuration {
	return BreakDuration{
		choices:  []string{"5", "10", "15", "20"},
		cursor:   0,
		selected: "",
	}
}

type SessionCount struct {
	choices  []string
	cursor   int
	selected string
}

func NewSessionCount() SessionCount {
	return SessionCount{
		choices:  []string{"4", "5", "6", "7"},
		cursor:   0,
		selected: "",
	}
}

type KeyMap struct {
	Start   key.Binding
	Stop    key.Binding
	Up      key.Binding
	Down    key.Binding
	Enter   key.Binding
	Init    key.Binding
	Confirm key.Binding
	Reset   key.Binding
	Quit    key.Binding
}

func NewKeyMap() KeyMap {
    km := KeyMap{
		Start: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "start"),
		),
		Stop: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "stop"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("k or up", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("j or down", "down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter", " "),
			key.WithHelp("enter", "enter"),
		),
		Init: key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "init"),
		),
		Confirm: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "confirm"),
		),
		Reset: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "reset"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}

    km.Start.SetEnabled(false)
    km.Stop.SetEnabled(false)
    km.Reset.SetEnabled(false)

    return km
}

type Model struct {
	KeyMap          KeyMap
	Help            help.Model
	WorkingDuration WorkingDuration
	BreakDuration   BreakDuration
	SessionCount    SessionCount

	State              string
	CurrentWorkSession int
	Timer              timer.Model
	TimerInitialized   bool
}

func New() Model {
	m := Model{
		KeyMap:             NewKeyMap(),
		Help:               help.NewModel(),
		WorkingDuration:    NewWorkingDuration(),
		BreakDuration:      NewBreakDuration(),
		SessionCount:       NewSessionCount(),

		State:              state.ChooseWorkingDuration,
		CurrentWorkSession: 0,
		TimerInitialized:   false,
	}
	return m
}

func (m Model) HasSelectedWorkingDuration() bool {
	return m.WorkingDuration.selected != ""
}

func (m Model) HasSelectedBreakDuration() bool {
	return m.BreakDuration.selected != ""
}

func (m Model) HasSelectedSessionCount() bool {
	return m.SessionCount.selected != ""
}

func (m Model) CurrentCursor() int {
	switch m.State {
	case state.ChooseWorkingDuration:
		return m.WorkingDuration.cursor
	case state.ChooseBreakDuration:
		return m.BreakDuration.cursor
	case state.ChooseSessionCount:
		return m.SessionCount.cursor
	}

	return 0
}

func (m Model) CurrentSelectedChoice() string {
	switch m.State {
	case state.ChooseWorkingDuration:
		return m.WorkingDuration.selected
	case state.ChooseBreakDuration:
		return m.BreakDuration.selected
	case state.ChooseSessionCount:
		return m.SessionCount.selected
	}

	return ""
}

func (m Model) CurrentChoices() []string {
	switch m.State {
	case state.ChooseWorkingDuration:
		return m.WorkingDuration.choices
	case state.ChooseBreakDuration:
		return m.BreakDuration.choices
	case state.ChooseSessionCount:
		return m.SessionCount.choices
	}

	return []string{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return HandleUpdate(msg, m)
}

func (m Model) View() string {
	return CreateView(m)
}
