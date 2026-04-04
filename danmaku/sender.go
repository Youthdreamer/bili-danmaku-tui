package danmaku

import (
	"errors"
	"regexp"

	"github.com/Akegarasu/blivedm-go/api"
)

func parseBiliVerify(cookie string) (*api.BiliVerify, error) {
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

func UserSendDanmaku(content, roomID, cookie string) error {
	verify, err := parseBiliVerify(cookie)
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
