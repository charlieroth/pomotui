package main

import (
	"fmt"
	"os"

	"github.com/charlieroth/pomotui/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := model.New()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
