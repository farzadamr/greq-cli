package ui

import "github.com/charmbracelet/lipgloss"

var (
	Theme = struct {
		Title      lipgloss.Style
		Error      lipgloss.Style
		Success    lipgloss.Style
		PurpleText lipgloss.Style
	}{
		Title:      lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00A7F7")),
		Error:      lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF4D4D")),
		Success:    lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#33CC66")),
		PurpleText: lipgloss.NewStyle().Foreground(lipgloss.Color("#ce08ffff")),
	}

	Line = "________________________\n"
)
