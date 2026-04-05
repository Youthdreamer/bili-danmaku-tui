package danmaku

import (
	"fmt"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
)

func Start(roomID int, cookie string, onMsg func(string)) (err error) {

	c := client.NewClient(roomID)
	if cookie == "" {
		fmt.Println("⚠️ 未设置 BLIVE_COOKIE，将使用匿名模式")
	} else {
		c.SetCookie(cookie)
	}

	c.OnDanmaku(func(d *message.Danmaku) {
		text := fmt.Sprintf("-> [%s] %s", d.Sender.Uname, d.Content)

		// 把弹幕发给 TUI
		onMsg(text)
	})

	if err := c.Start(); err != nil {
		// 如果连接过程中崩溃或断开，直接退出
		fmt.Printf("\n❌ 弹幕服务已停止: %v\n", err)
	}

	return nil
}
