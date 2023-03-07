package model

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) HandleUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		return m.HandleTimerTickMsg(msg)
	case timer.StartStopMsg:
		return m.HandleTimerStartStopMsg(msg)
	case timer.TimeoutMsg:
		return m.HandleTimerTimeout()
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Quit):
			return m.HandleQuit()
		case key.Matches(msg, m.KeyMap.Up):
			return m.HandleUp()
		case key.Matches(msg, m.KeyMap.Down):
			return m.HandleDown()
		case key.Matches(msg, m.KeyMap.Enter):
			return m.HandleEnter()
		case key.Matches(msg, m.KeyMap.Confirm):
			return m.HandleConfirm()
		case key.Matches(msg, m.KeyMap.Continue):
			return m.HandleContinue()
		case key.Matches(msg, m.KeyMap.Start, m.KeyMap.Stop):
			return m.HandleStartStop()
		}
	}

	return m, nil
}
