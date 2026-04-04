package tui

import (
	"runtime"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/mattn/go-runewidth"
)

// 定义跨平台换行符
var EOL = func() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}()

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

	wrappedLines := make([]string, 0, len(m.Lines))
	for _, line := range m.Lines {
		wrappedLines = append(wrappedLines, wrapLine(line, m.Width)...)
	}

	availableHeight := m.Height - 3
	if availableHeight < 0 {
		availableHeight = 0
	}

	if len(wrappedLines) > availableHeight {
		wrappedLines = wrappedLines[len(wrappedLines)-availableHeight:]
	}

	danmuku := strings.Join(wrappedLines, EOL)
	lineCount := len(wrappedLines)
	freeSpace := availableHeight - lineCount

	padding := ""
	if freeSpace > 0 {
		padding = strings.Repeat(EOL, freeSpace)
	}

	separator := EOL + EOL + EOL

	InputView := m.Input.View()
	v := tea.NewView(danmuku + padding + separator + InputView)

	return v
}

func wrapLine(line string, width int) []string {
	if width <= 0 {
		return nil
	}
	if line == "" {
		return []string{""}
	}

	lines := make([]string, 0, 1)
	var builder strings.Builder
	currentWidth := 0

	for _, r := range line {
		if r == '\n' {
			lines = append(lines, builder.String())
			builder.Reset()
			currentWidth = 0
			continue
		}

		runeWidth := runewidth.RuneWidth(r)
		if currentWidth+runeWidth > width && builder.Len() > 0 {
			lines = append(lines, builder.String())
			builder.Reset()
			currentWidth = 0
		}

		builder.WriteRune(r)
		currentWidth += runeWidth
	}

	lines = append(lines, builder.String())
	return lines
}
