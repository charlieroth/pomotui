package model

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		return HandleTimerTickMsg(m, msg)
	case timer.StartStopMsg:
		return HandleTimerStartStopMsg(m, msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Quit):
			return HandleQuit(m)
		case key.Matches(msg, m.KeyMap.Up):
			return HandleUp(m)
		case key.Matches(msg, m.KeyMap.Down):
			return HandleDown(m)
		case key.Matches(msg, m.KeyMap.Enter):
			return HandleEnter(m)
		case key.Matches(msg, m.KeyMap.Confirm):
			return HandleConfirm(m)
		case key.Matches(msg, m.KeyMap.Start, m.KeyMap.Stop):
            return HandleStartStop(m)
		}
	}

	return m, nil
}
