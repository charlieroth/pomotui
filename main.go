package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
    "github.com/charlieroth/pomotui/model"
)

func main() {
    model := model.New()
	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
