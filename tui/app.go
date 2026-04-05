package tui

import (
	"fmt"
	"os"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"github.com/Akegarasu/blivedm-go/api"
	"github.com/Youthdreamer/bili-danmaku-tui/danmaku"
)

func Run(roomID string, cookie string) error {

	// 预检：先验证房间是否有效
	id, err := strconv.Atoi(roomID)
	roomInfo, err := api.GetRoomInfo(id)
	if err != nil || roomInfo.Code != 0 {
		return fmt.Errorf("房间 %s 无效或不存在", roomID)
	}

	// 预检通过后，再创建客户端
	m := NewModel(roomID, cookie)
	p := tea.NewProgram(m)

	// 👉 启动弹幕监听
	go func() {
		err := danmaku.Start(id, cookie, func(msg string) {
			p.Send(DanmakuMsg(msg))
		})
		if err != nil {
			fmt.Println("弹幕服务错误:", err)
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("❌ 界面运行出错: %v\n", err)
		os.Exit(1)
	}

	return nil
}
