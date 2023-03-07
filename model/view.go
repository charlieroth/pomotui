package model

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charlieroth/pomotui/state"
	"github.com/charlieroth/pomotui/ui"
	"github.com/charmbracelet/bubbles/key"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/gen2brain/beeep"
)

type soundInfo struct {
	format   beep.Format
	streamer beep.StreamSeeker
	done     chan bool
}

/*
	Previous model state is saved in view in order to

catch state changes like an interval ending
*/
var previousState string

func decodeSound() soundInfo {
	var (
		err   error
		sound soundInfo
	)
	data, err := os.Open("resources/ring_sound.mp3")
	if err != nil {
		log.Panicf("Error opening sound file: %v", err)
	}
	defer data.Close()

	sound.streamer, sound.format, err = mp3.Decode(data)
	if err != nil {
		log.Panicf("Error decoding sound file: %v", err)
	}
	return sound
}

func init() {
	sound := decodeSound()
	err := speaker.Init(sound.format.SampleRate, sound.format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatalf("Error initializing speaker: %v", err)
	}

	sound.done = make(chan bool)
}

func playRingSound(sound soundInfo) {
	speaker.Lock()
	defer speaker.Unlock()
	speaker.Clear()
	speaker.Play(beep.Seq(sound.streamer, beep.Callback(func() {
		sound.done <- true
	})))
	<-sound.done
	err := sound.streamer.Seek(0)
	if err != nil {
		log.Panicf("Error seeking through sound file: %v", err)
	}
}

func CreateView(m Model) string {

	view := GetTitle(m)

	switch m.State {
	case state.ChooseWorkingDuration, state.ChooseBreakDuration, state.ChooseLongBreakDuration, state.ChooseSessionCount:
		view += ChoicesView(m)
	case state.Working, state.Break, state.LongBreak:
		view += MainView(m)
	}
	view += HelpView(m)

	sound := decodeSound()

	if breakEndJustHappened(m) {
		err := beeep.Notify("End of break", "C'mon, back to work", "")
		if err != nil {
			log.Panicf("Error showing notification: %v", err)
			return ""
		}
		go playRingSound(sound)
	} else if breakJustHappened(m) {
		err := beeep.Notify("Work inteval finished", "Time for a break!", "")
		if err != nil {
			log.Panicf("Error showing notification: %v", err)
			return ""
		}
		go playRingSound(sound)
	}
	previousState = m.State
	return view
}

func breakEndJustHappened(m Model) bool {
	return (m.State == state.Working &&
		(previousState == state.Break || previousState == state.LongBreak))
}

func breakJustHappened(m Model) bool {
	return (previousState == state.Working &&
		(m.State == state.Break || m.State == state.LongBreak))
}

func WorkingDurationTitle() string {
	return "Work Duration:\n"
}

func BreakDurationTitle() string {
	return "Break Duration:\n"
}

func LongBreakDurationTitle() string {
	return "Long Break Duration:\n"
}
func SessionCountTitle() string {
	return "Sesson Count:\n"
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
		panic("failed to convert session count from string to int")
	}

	var s strings.Builder

	for i := 1; i <= sessionCount; i++ {
		if m.CurrentWorkSession >= i {
			s.WriteString(" " + ui.ActiveString("•"))
		} else {
			s.WriteString(" " + ui.InactivateString("•"))
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
		m.KeyMap.Continue,
		m.KeyMap.Reset,
		m.KeyMap.Quit,
	})
}
