package ui

import "github.com/charmbracelet/lipgloss"

func InactivateString(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render(s)
}

func ActiveString(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render(s)
}
