package tui

import (
	"runtime"
	"strings"

	tea "charm.land/bubbletea/v2"
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

	// 1. 处理弹幕截断：如果单行太长，显示 >>
	processedLines := make([]string, len(m.Lines))
	suffix := " >>"
	// 预留一些宽度给可能的滚动条或边距（建议减去 2 到 4 个单位）
	limit := m.Width - len(suffix)

	for i, line := range m.Lines {
		// 注意：如果包含中文，len(line) 不准确，建议后续使用 utf8.RuneCountInString
		if len(line) > m.Width {
			if limit > 0 {
				processedLines[i] = line[:limit] + suffix
			} else {
				processedLines[i] = line[:m.Width]
			}
		} else {
			processedLines[i] = line
		}
	}

	// 2. 组合内容
	danmuku := strings.Join(processedLines, EOL)
	lineCount := len(processedLines)

	// 3. 重新计算 padding 逻辑
	// 视觉空 2 行需要 3 个换行符：
	// 第 1 个换行结束弹幕区，第 2、3 个换行产生两个空行
	// 所以总占用高度是：lineCount + 2 (空行) + 1 (输入框) = lineCount + 3
	freeSpace := m.Height - lineCount - 3

	padding := ""
	if freeSpace > 0 {
		padding = strings.Repeat(EOL, freeSpace)
	}

	// 4. 组装：弹幕 + 填充 + 3个换行符(实现空两行) + 输入框
	// 这里直接写 3 个 EOL，确保无论 padding 是否存在，间隔都固定
	separator := EOL + EOL + EOL

	InputView := m.Input.View()
	v := tea.NewView(danmuku + padding + separator + InputView)

	return v
}
