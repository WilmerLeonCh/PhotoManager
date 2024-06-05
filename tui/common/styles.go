package common

import "github.com/charmbracelet/lipgloss"

// https://colorhunt.co/palette/cde8e5eef7ff7ab2b24d869c
const (
	primaryColor   = lipgloss.Color("#4D869C")
	secondaryColor = lipgloss.Color("#7AB2B2")
	thirdColor     = lipgloss.Color("#CDE8E5")
	grayColor      = lipgloss.Color("#959a9c")
)

var (
	TableStyle = lipgloss.
			NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(primaryColor).
			Padding(0, 1)
	HeaderTableStyle = lipgloss.
				NewStyle().
				Foreground(secondaryColor).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(thirdColor).
				BorderBottom(true).
				Bold(true)
	SelectedTableStyle = lipgloss.
				NewStyle().
				Foreground(primaryColor).
				Background(thirdColor)
	TitleStyle = lipgloss.
			NewStyle().
			Foreground(primaryColor).
			Bold(true)
	PromptStyle = lipgloss.
			NewStyle().
			Foreground(secondaryColor).
			Bold(true).
			Margin(1, 0, 1, 0)
	PlaceholderStyle = lipgloss.
				NewStyle().
				Foreground(grayColor).
				Padding(0, 1)
	OptionListStyle = lipgloss.
			NewStyle().
			Foreground(secondaryColor).
			Bold(true).
			Margin(1, 0, 1, 0)
)
