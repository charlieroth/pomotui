package main

import (
	"fmt"
	"github.com/charlieroth/pomotui/model"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	model := model.New()
	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
