package template

import "github.com/line/line-bot-sdk-go/linebot"

type Buttons struct {
	imagePath string
	Title     string
	subTitle  string
	button    []linebot.PostbackTemplateAction
}

func buttonsTemplate(btn Buttons) *linebot.TemplateMessage {
	t := linebot.NewButtonsTemplate(btn.imagePath, btn.Title, btn.subTitle, btn.button...)
	msg := linebot.NewTemplateMessage("buttons template", t)
	return msg
}
