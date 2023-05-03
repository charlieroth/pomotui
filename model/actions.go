package model

import (
	"strconv"
	"time"

	"github.com/charlieroth/pomotui/state"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelHandler interface {
	HandleQuit(tea.Model, tea.Cmd)
	HandleStartStop(tea.Model, tea.Cmd)
	HandleUp(tea.Model, tea.Cmd)
	HandleDown(tea.Model, tea.Cmd)
	HandleConfirm(tea.Model, tea.Cmd)
	HandleContinue(tea.Model, tea.Cmd)
	HandleEnter(tea.Model, tea.Cmd)
	HandleTimerTickMsg(timer.TickMsg) (tea.Model, tea.Cmd)
	HandleTimerStartStopMsg(timer.StartStopMsg) (tea.Model, tea.Cmd)
	HandleTimerTimeout() (tea.Model, tea.Cmd)
	HandleUpdate(msg tea.Msg) (tea.Model, tea.Cmd)
}

func (m *Model) HandleQuit() (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m *Model) HandleStartStop() (tea.Model, tea.Cmd) {
	if !m.TimerInitialized {
		m.TimerInitialized = true
		m.KeyMap.Stop.SetEnabled(true)
		m.KeyMap.Start.SetEnabled(false)
		return m, m.Timer.Init()
	}

	return m, m.Timer.Toggle()
}

func (m *Model) HandleUp() (tea.Model, tea.Cmd) {
	switch m.State {
	case state.ChooseWorkingDuration:
		if m.WorkingDuration.cursor > 0 {
			m.WorkingDuration.cursor--
		}
		return m, nil
	case state.ChooseBreakDuration:
		if m.BreakDuration.cursor > 0 {
			m.BreakDuration.cursor--
		}
		return m, nil
	case state.ChooseLongBreakDuration:
		if m.LongBreakDuration.cursor > 0 {
			m.LongBreakDuration.cursor--
		}
		return m, nil
	case state.ChooseSessionCount:
		if m.SessionCount.cursor > 0 {
			m.SessionCount.cursor--
		}
		return m, nil
	}

	return m, nil
}

func (m *Model) HandleDown() (tea.Model, tea.Cmd) {
	switch m.State {
	case state.ChooseWorkingDuration:
		if m.WorkingDuration.cursor < len(m.WorkingDuration.choices)-1 {
			m.WorkingDuration.cursor++
		}
		return m, nil
	case state.ChooseBreakDuration:
		if m.BreakDuration.cursor < len(m.BreakDuration.choices)-1 {
			m.BreakDuration.cursor++
		}
		return m, nil
	case state.ChooseLongBreakDuration:
		if m.LongBreakDuration.cursor < len(m.LongBreakDuration.choices)-1 {
			m.LongBreakDuration.cursor++
		}
		return m, nil
	case state.ChooseSessionCount:
		if m.SessionCount.cursor < len(m.SessionCount.choices)-1 {
			m.SessionCount.cursor++
		}
		return m, nil
	}

	return m, nil
}

func (m *Model) HandleConfirm() (tea.Model, tea.Cmd) {
	switch m.State {
	case state.ChooseWorkingDuration:
		m.KeyMap.Start.SetEnabled(false)
		m.KeyMap.Stop.SetEnabled(false)
		m.KeyMap.Reset.SetEnabled(false)
		if !m.HasSelectedWorkingDuration() {
			return m, nil
		}

		m.State = state.ChooseBreakDuration
	case state.ChooseBreakDuration:
		m.KeyMap.Start.SetEnabled(false)
		m.KeyMap.Stop.SetEnabled(false)
		m.KeyMap.Reset.SetEnabled(false)
		if !m.HasSelectedBreakDuration() {
			return m, nil
		}

		m.State = state.ChooseLongBreakDuration
	case state.ChooseLongBreakDuration:
		m.KeyMap.Start.SetEnabled(false)
		m.KeyMap.Stop.SetEnabled(false)
		m.KeyMap.Reset.SetEnabled(false)
		if !m.HasSelectLongBreakDuration() {
			return m, nil
		}

		m.State = state.ChooseSessionCount
	case state.ChooseSessionCount:
		if !m.HasSelectedSessionCount() {
			return m, nil
		}

		// transition into "working" state & set first working session timer
		m.CurrentWorkSession = 1
		m.State = state.Working
		selectedTime, err := strconv.Atoi(m.WorkingDuration.selected)
		if err != nil {
			panic("Failed to convert working duration time to int")
		}

		amountOfTime := time.Duration(selectedTime) * time.Minute
		m.Timer = timer.NewWithInterval(amountOfTime, time.Second)
		m.KeyMap.Start.SetEnabled(true)
		m.KeyMap.Stop.SetEnabled(true)
		m.KeyMap.Up.SetEnabled(false)
		m.KeyMap.Down.SetEnabled(false)
		m.KeyMap.Enter.SetEnabled(false)
		m.KeyMap.Confirm.SetEnabled(false)
	}
	return m, nil
}

func (m *Model) HandleContinue() (tea.Model, tea.Cmd) {
	if m.State == state.Break {
		workTimeInt, err := strconv.Atoi(m.WorkingDuration.selected)
		if err != nil {
			panic("Failed to convert work duration time to int")
		}

		workTime := time.Duration(workTimeInt) * time.Minute
		m.Timer = timer.NewWithInterval(workTime, time.Second)
		m.TimerInitialized = false

		m.KeyMap.Stop.SetEnabled(false)
		m.KeyMap.Start.SetEnabled(true)
		m.KeyMap.Continue.SetEnabled(false)

		m.CurrentWorkSession++
		m.State = state.Working

		return m, nil
	}
	return nil, nil
}

func (m *Model) HandleEnter() (tea.Model, tea.Cmd) {
	switch m.State {
	case state.ChooseWorkingDuration:
		if m.WorkingDuration.selected == "" {
			m.WorkingDuration.selected = m.WorkingDuration.choices[m.WorkingDuration.cursor]
			return m, nil
		}

		if m.WorkingDuration.choices[m.WorkingDuration.cursor] != m.WorkingDuration.selected {
			m.WorkingDuration.selected = m.WorkingDuration.choices[m.WorkingDuration.cursor]
			return m, nil
		}

		m.WorkingDuration.selected = ""
		return m, nil
	case state.ChooseBreakDuration:
		if m.BreakDuration.selected == "" {
			m.BreakDuration.selected = m.BreakDuration.choices[m.BreakDuration.cursor]
			return m, nil
		}

		if m.BreakDuration.choices[m.BreakDuration.cursor] != m.BreakDuration.selected {
			m.BreakDuration.selected = m.BreakDuration.choices[m.BreakDuration.cursor]
			return m, nil
		}

		m.BreakDuration.selected = ""
		return m, nil
	case state.ChooseLongBreakDuration:
		if m.LongBreakDuration.selected == "" {
			m.LongBreakDuration.selected = m.LongBreakDuration.choices[m.LongBreakDuration.cursor]
			return m, nil
		}

		if m.LongBreakDuration.choices[m.LongBreakDuration.cursor] != m.LongBreakDuration.selected {
			m.LongBreakDuration.selected = m.LongBreakDuration.choices[m.LongBreakDuration.cursor]
			return m, nil
		}

		m.LongBreakDuration.selected = ""
		return m, nil
	case state.ChooseSessionCount:
		if m.SessionCount.selected == "" {
			m.SessionCount.selected = m.SessionCount.choices[m.SessionCount.cursor]
			return m, nil
		}
		if m.SessionCount.choices[m.SessionCount.cursor] != m.SessionCount.selected {
			m.SessionCount.selected = m.SessionCount.choices[m.SessionCount.cursor]
		}

		m.SessionCount.selected = ""
		return m, nil
	}

	return m, nil
}

func (m *Model) HandleTimerTickMsg(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Timer, cmd = m.Timer.Update(msg)
	return m, cmd
}

func (m *Model) HandleTimerStartStopMsg(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Timer, cmd = m.Timer.Update(msg)
	m.KeyMap.Stop.SetEnabled(m.Timer.Running())
	m.KeyMap.Start.SetEnabled(!m.Timer.Running())
	return m, cmd
}

func (m *Model) HandleTimerTimeout() (tea.Model, tea.Cmd) {
	sessionCount, err := strconv.Atoi(m.SessionCount.selected)
	if err != nil {
		panic("Failed to convert work duration time to int")
	}

	// completed last working session, transition to long break
	if m.State == state.Working && m.CurrentWorkSession == sessionCount {
		breakTimeInt, err := strconv.Atoi(m.LongBreakDuration.selected)
		if err != nil {
			panic("Failed to convert long break duration time to int")
		}

		breakTime := time.Duration(breakTimeInt) * time.Minute
		m.Timer = timer.NewWithInterval(breakTime, time.Second)
		m.TimerInitialized = false

		m.KeyMap.Stop.SetEnabled(false)
		m.KeyMap.Start.SetEnabled(true)
		m.KeyMap.Continue.SetEnabled(true)

		m.CurrentWorkSession = 0
		m.State = state.LongBreak
		return m, nil
	}

	// completed 1 of X working sessions, transition to break
	if m.State == state.Working {
		breakTimeInt, err := strconv.Atoi(m.BreakDuration.selected)
		if err != nil {
			panic("Failed to convert break duration time to int")
		}

		breakTime := time.Duration(breakTimeInt) * time.Minute
		m.Timer = timer.NewWithInterval(breakTime, time.Second)
		m.TimerInitialized = false

		m.KeyMap.Stop.SetEnabled(false)
		m.KeyMap.Start.SetEnabled(true)
		m.KeyMap.Continue.SetEnabled(true)

		m.State = state.Break
		return m, nil
	}

	// completed 1 of X breaks, transition to working
	if m.State == state.Break {
		workTimeInt, err := strconv.Atoi(m.WorkingDuration.selected)
		if err != nil {
			panic("Failed to convert work duration time to int")
		}

		workTime := time.Duration(workTimeInt) * time.Minute
		m.Timer = timer.NewWithInterval(workTime, time.Second)
		m.TimerInitialized = false

		m.KeyMap.Stop.SetEnabled(false)
		m.KeyMap.Start.SetEnabled(true)
		m.KeyMap.Continue.SetEnabled(false)

		m.CurrentWorkSession++
		m.State = state.Working

		return m, nil
	}

	// completed long break, transition to working
	if m.State == state.LongBreak {
		workTimeInt, err := strconv.Atoi(m.WorkingDuration.selected)
		if err != nil {
			panic("Failed to convert working duration time to int")
		}

		workTime := time.Duration(workTimeInt) * time.Minute
		m.Timer = timer.NewWithInterval(workTime, time.Second)
		m.TimerInitialized = false

		m.KeyMap.Stop.SetEnabled(false)
		m.KeyMap.Start.SetEnabled(true)

		m.CurrentWorkSession = 1
		m.State = state.Working
		return m, nil
	}

	return m, nil
}
