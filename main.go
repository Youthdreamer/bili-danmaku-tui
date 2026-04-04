package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/Akegarasu/blivedm-go/api"
	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
	"github.com/joho/godotenv"
)

type DanmakuMsg string

type model struct {
	width  int
	height int
	input  textinput.Model
	lines  []string
	roomID string // 添加房间ID
	cookie string // 添加cookie
}

func (m model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.Code {
		case tea.KeyEnter:
			content := m.input.Value()
			if content != "" {
				m.input.Reset()
				return m, sendDanmakuCmd(content, m.roomID, m.cookie)
			}
		default:
			switch msg.String() {
			case "ctrl+c", "q", "esc":
				return m, tea.Quit
			case "ctrl+l":
				m.lines = nil
				return m, tea.ClearScreen
			}
		}

	case DanmakuMsg:
		m.lines = append(m.lines, string(msg))
		if len(m.lines) > 20 {
			m.lines = m.lines[len(m.lines)-20:]
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	m.input, cmd = m.input.Update(msg)

	return m, cmd
}

func (m model) View() tea.View {
	if m.width == 0 {
		m.width = 80
	}
	if m.height == 0 {
		m.height = 24
	}

	danmuku := strings.Join(m.lines, "\n")
	lineCount := len(m.lines)
	padding := ""
	if m.height-lineCount-1 > 0 {
		padding = strings.Repeat("\n", m.height-lineCount-1)
	}
	inputView := m.input.View()

	v := tea.NewView(danmuku + padding + inputView)

	return v
}

func ParseBiliVerify(cookie string) (*api.BiliVerify, error) {
	verify := &api.BiliVerify{}

	// 解析 bili_jct (CSRF token)
	reCsrf := regexp.MustCompile(`bili_jct=([^;]+)`)
	result := reCsrf.FindStringSubmatch(cookie)
	if len(result) > 1 {
		verify.Csrf = result[1]
	}

	// 解析 SESSDATA
	reSess := regexp.MustCompile(`SESSDATA=([^;]+)`)
	result = reSess.FindStringSubmatch(cookie)
	if len(result) > 1 {
		verify.SessData = result[1]
	}

	if verify.Csrf == "" || verify.SessData == "" {
		return nil, errors.New("cookie中缺少必要的bili_jct或SESSDATA字段")
	}

	return verify, nil
}

func sendDanmakuCmd(content, roomID, cookie string) tea.Cmd {
	return func() tea.Msg {

		verify, err := ParseBiliVerify(cookie)
		if err != nil {
			return err
		}
		dmReq := &api.DanmakuRequest{
			Msg:      content,
			RoomID:   roomID,
			Bubble:   "0",
			Color:    "16777215",
			FontSize: "25",
			Mode:     "1",
			DmType:   "0",
		}
		_, err = api.SendDanmaku(dmReq, &api.BiliVerify{
			Csrf:     verify.Csrf,
			SessData: verify.SessData,
		})
		if err != nil {
			return err
		}
		return nil
	}
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
	// 获取Cookie
	cookie := os.Getenv("BLIVE_COOKIE")

	ti := textinput.New()
	ti.Placeholder = "在此输入弹幕..."
	ti.Focus() // 👈 非常重要：如果不 Focus，它会忽略所有按键
	ti.Placeholder = "输入弹幕内容，按 Enter 发送..."
	ti.SetWidth(50)

	// 初始化 model 时传入 roomID 和 cookie
	initialModel := model{
		roomID: roomID,
		cookie: cookie,
		height: 24, // 默认高度
		input:  ti,
	}

	p := tea.NewProgram(initialModel)

	// 👉 启动弹幕监听（放 goroutine）
	go func() {
		id, _ := strconv.Atoi(roomID)
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
