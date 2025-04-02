package cli

import "github.com/fatih/color"

// ColorScheme define o esquema de cores usado na interface
type ColorScheme struct {
	Title     *color.Color
	Subtitle  *color.Color
	Success   *color.Color
	Error     *color.Color
	Warning   *color.Color
	Info      *color.Color
	Prompt    *color.Color
	Highlight *color.Color
}

// NewColorScheme cria um novo esquema de cores
func NewColorScheme() *ColorScheme {
	return &ColorScheme{
		Title:     color.New(color.FgHiCyan, color.Bold),
		Subtitle:  color.New(color.FgCyan),
		Success:   color.New(color.FgHiGreen, color.Bold),
		Error:     color.New(color.FgHiRed, color.Bold),
		Warning:   color.New(color.FgHiYellow),
		Info:      color.New(color.FgHiBlue),
		Prompt:    color.New(color.FgHiMagenta),
		Highlight: color.New(color.FgHiWhite, color.Bold),
	}
}
