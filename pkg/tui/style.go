package tui

import "github.com/charmbracelet/lipgloss"

var blue = lipgloss.Color("#6991ac")

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(blue).
	Padding(2, 4, 2, 4).
	Margin(2, 4, 2, 4)
