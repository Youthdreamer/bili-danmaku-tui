package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Youthdreamer/bili-danmaku-tui/config"
	"github.com/Youthdreamer/bili-danmaku-tui/danmaku"
	"github.com/Youthdreamer/bili-danmaku-tui/tui"
)

func main() {
	// 检查是否输入了参数
	if len(os.Args) < 2 {
		fmt.Println("请提供直播间ID")
		return
	}

	roomID := os.Args[1]
	cookie := config.Load()

	m := tui.NewModel(roomID, cookie)
	p := tea.NewProgram(m)

	// 👉 启动弹幕监听
	go func() {
		err := danmaku.Start(roomID, cookie, func(msg string) {
			p.Send(tui.DanmakuMsg(msg))
		})
		if err != nil {
			fmt.Println("弹幕服务错误:", err)
			os.Exit(1)
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("❌ 界面运行出错: %v\n", err)
		os.Exit(1)
	}
}
