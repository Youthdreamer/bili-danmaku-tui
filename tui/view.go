package tui

import (
	"fmt"
	"runtime"
	"strings"
	"unicode/utf8"

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

	availableHeight := max(0, m.Height-3)

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

	inputView := m.Input.View()
	currentLen := utf8.RuneCountInString(m.Input.Value())
	counterStr := fmt.Sprintf(" [%d/%d]", currentLen, maxLen)

	// 2. 动态位置逻辑：计算留给输入框的“安全空间”
	// 假设我们希望输入框至少能显示 10 个字符宽
	const minInputSpace = 10

	indent := m.Width - runewidth.StringWidth(counterStr)
	hintLine := ""
	if indent > 0 {
		hintLine = strings.Repeat(" ", max(0, indent)) + counterStr
	} else {
		hintLine = counterStr
	}

	v := tea.NewView(danmuku + padding + EOL + hintLine + EOL + inputView)
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
