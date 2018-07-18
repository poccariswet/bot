package template

import "github.com/line/line-bot-sdk-go/linebot"

type Confirms struct {
	Title string
	Left  linebot.TemplateAction
	Right linebot.TemplateAction
}

func ConfirmsTemplate(c Confirms) *linebot.TemplateMessage {
	return linebot.NewTemplateMessage("confirms template",
		&linebot.ConfirmTemplate{
			Text:    c.Title,
			Actions: []linebot.TemplateAction{c.Left, c.Right},
		})
}
