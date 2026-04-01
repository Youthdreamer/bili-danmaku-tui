package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

type DanmakuMsg string

type model struct {
	lines []string
}

func (m model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case DanmakuMsg:
		m.lines = append(m.lines, string(msg))

		// 只保留最近 20 条
		if len(m.lines) > 20 {
			m.lines = m.lines[len(m.lines)-20:]
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "c":
			m.lines = nil
			return m, tea.ClearScreen
		}
	}

	return m, nil
}

func (m model) View() string {
	return strings.Join(m.lines, "\n")
}

func main() {
	godotenv.Load()
	// 检查是否输入了参数
	if len(os.Args) < 2 {
		fmt.Println("请提供直播间ID，例如: go run main.go 22222")
		return
	}

	// 获取第一个参数
	roomID := os.Args[1]
	id, _ := strconv.Atoi(roomID)

	// 获取Cookie
	cookie := os.Getenv("BLIVE_COOKIE")

	p := tea.NewProgram(model{})

	// 👉 启动弹幕监听（放 goroutine）
	go func() {
		c := client.NewClient(id)
		if cookie == "" {
			fmt.Println("⚠️ 未设置 BLIVE_COOKIE，将使用匿名模式")
		} else {
			c.SetCookie(cookie)
		}

		c.OnDanmaku(func(d *message.Danmaku) {
			text := fmt.Sprintf("[%s] %s", d.Sender.Uname, d.Content)

			// 👉 把弹幕发给 TUI
			p.Send(DanmakuMsg(text))
		})

		if err := c.Start(); err != nil {
			// 如果连接过程中崩溃或断开，直接退出
			fmt.Printf("\n❌ 弹幕服务已停止: %v\n", err)
			os.Exit(1)
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("❌ 界面运行出错: %v\n", err)
		os.Exit(1)
	}
}
