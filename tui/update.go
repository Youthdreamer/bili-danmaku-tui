package tui

import (
	"unicode/utf8"

	tea "charm.land/bubbletea/v2"
	"github.com/Youthdreamer/bili-danmaku-tui/danmaku"
)

// b站弹幕上限为 40 字符
const maxLen = 40

// 定义发送弹幕的 Cmd
func SendDanmakuCmd(content, roomID, cookie string) tea.Cmd {
	return func() tea.Msg {
		if err := danmaku.UserSendDanmaku(content, roomID, cookie); err != nil {
			return DanmakuMsg("❌ 发送失败: " + err.Error())
		}
		return nil
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// 获取当前输入的字符数 (rune count)
	currentLen := utf8.RuneCountInString(m.Input.Value())

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		// --- 1. 输入拦截逻辑 ---
		// 如果字数已满，且按下的不是功能键（回车、退格、Esc等），则直接拦截，不更新状态
		if currentLen >= maxLen &&
			msg.Code != tea.KeyEnter &&
			msg.Code != tea.KeyBackspace &&
			msg.Code != tea.KeyDelete &&
			msg.Code != tea.KeyEsc {
			return m, nil
		}
		switch msg.Code {
		case tea.KeyEnter:
			content := m.Input.Value()
			if content != "" {
				m.Input.Reset()
				return m, SendDanmakuCmd(content, m.RoomID, m.Cookie)
			}
		default:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "ctrl+l":
				m.Lines = nil
				return m, tea.ClearScreen
			}
		}

	case DanmakuMsg:
		m.Lines = append(m.Lines, string(msg))
		if len(m.Lines) > 20 {
			m.Lines = m.Lines[len(m.Lines)-20:]
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}

	m.Input, cmd = m.Input.Update(msg)

	return m, cmd
}
