package tui

import "charm.land/bubbles/v2/textinput"

type DanmakuMsg string

type Model struct {
	Width, Height int
	Input         textinput.Model
	Lines         []string
	RoomID        string
	Cookie        string
}

func NewModel(roomID, cookie string) Model {
	ti := textinput.New()
	ti.Placeholder = "输入弹幕，回车发送..."
	ti.Focus()
	ti.SetWidth(50)

	return Model{
		Input:  ti,
		RoomID: roomID,
		Cookie: cookie,
	}
}
