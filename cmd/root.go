package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/Youthdreamer/bili-danmaku-tui/config"
	"github.com/Youthdreamer/bili-danmaku-tui/tui"
	"github.com/spf13/cobra"
)

var (
	roomID    int
	Version   = "dev"
	GitCommit = "none"
	BuildTime = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "bili-danmaku-tui",
	Short: "Bilibili直播终端弹幕助手",
	Long:  "一个用于发送和接收哔哩哔哩直播弹幕的终端UI客户端。",

	Version: Version,

	// 核心运行逻辑
	Run: func(cmd *cobra.Command, args []string) {
		var finalRoomID string

		// 优先级 1: 检查是否通过 -r 传入了参数
		if roomID != 0 {
			finalRoomID = strconv.Itoa(roomID)
		} else if len(args) > 0 {
			// 优先级 2: 检查是否通过第一个位置参数传入
			finalRoomID = args[0]
		}

		// 如果用户没传房间号
		if finalRoomID == "" {
			fmt.Println("❌ 错误: 请提供直播间 ID (例如: bili-danmaku-tui ID 或使用 -r ID)")
			cmd.Usage()
			return
		}

		// id := fmt.Sprintf("%d", finalRoomID)
		cookie := config.Load()

		if err := tui.Run(finalRoomID, cookie); err != nil {
			fmt.Printf("程序运行出错: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// flags
	rootCmd.Flags().IntVarP(&roomID, "room", "r", 0, "指定 B 站直播间 ID (必填)")
	rootCmd.SetVersionTemplate(getCustomVersion())

}

func getCustomVersion() string {
	return fmt.Sprintf(`Version:      %s
Git Commit:   %s
Build Time:   %s
Go Version:   %s
OS/Arch:      %s/%s
`, Version, GitCommit, BuildTime, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
