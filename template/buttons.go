package template

import "github.com/line/line-bot-sdk-go/linebot"

type Buttons struct {
	ImagePath string
	Title     string
	SubTitle  string
	Button    []linebot.TemplateAction
}

func ButtonsTemplate(btn Buttons) *linebot.TemplateMessage {
	//	t := linebot.NewButtonsTemplate(btn.ImagePath, btn.Title, btn.SubTitle, btn.Button...)
	msg := linebot.NewTemplateMessage("buttons template",
		&linebot.ButtonsTemplate{
			ThumbnailImageURL: btn.ImagePath,
			Title:             btn.Title,
			Text:              btn.SubTitle,
			Actions:           btn.Button,
		})
	return msg
}
