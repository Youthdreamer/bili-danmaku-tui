package tui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Init() tea.Cmd {
	return tea.ClearScreen
}
func (m Model) View() tea.View {
	if m.Width == 0 {
		m.Width = 80
	}
	if m.Height == 0 {
		m.Height = 24
	}

	danmuku := strings.Join(m.Lines, "\n")
	lineCount := len(m.Lines)
	padding := ""
	if m.Height-lineCount-1 > 0 {
		padding = strings.Repeat("\n", m.Height-lineCount-1)
	}
	InputView := m.Input.View()

	v := tea.NewView(danmuku + padding + InputView)

	return v
}
