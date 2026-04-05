package tui

import (
	"unicode/utf8"

	tea "charm.land/bubbletea/v2"
	"github.com/Youthdreamer/bili-danmaku-tui/danmaku"
)

type InitMsg string

func (m Model) Init() tea.Cmd {
	return tea.Sequence(
		func() tea.Msg { return tea.ClearScreen() },
		func() tea.Msg { return InitMsg("欢迎使用 Bilibili 弹幕助手！\n") },
	)
}

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

	switch msg := msg.(type) {
	case InitMsg:
		// 处理欢迎消息，可以保存到 model 中
		m.welcomeMsg = string(msg)

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
			case "ctrl+c":
				return m, tea.Quit
			case "ctrl+l":
				m.Lines = nil
				return m, tea.ClearScreen
			}
		}

	case DanmakuMsg:
		m.Lines = append(m.Lines, string(msg))
		if len(m.Lines) > 100 {
			m.Lines = m.Lines[len(m.Lines)-100:]
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height + 2
	}

	m.Input, cmd = m.Input.Update(msg)
	// 无论通过什么方式输入（英文、中文、粘贴），只要总长度超了，就立刻截断
	val := m.Input.Value()
	if utf8.RuneCountInString(val) > maxLen {
		// 将字符串转为 rune 切片进行安全截断，防止中文字符被切碎
		runes := []rune(val)
		m.Input.SetValue(string(runes[:maxLen]))

		// 3. 修正光标位置
		// 避免截断后光标停留在失效的索引位置导致渲染崩溃
		m.Input.SetCursor(maxLen)
	}
	return m, cmd
}
