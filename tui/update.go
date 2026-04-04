package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Youthdreamer/bili-danmaku-tui/danmaku"
)

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
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.Code {
		case tea.KeyEnter:
			content := m.Input.Value()
			if content != "" {
				m.Input.Reset()
				return m, SendDanmakuCmd(content, m.RoomID, m.Cookie)
			}
		default:
			switch msg.String() {
			case "ctrl+c", "q", "esc":
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
