package template

import "github.com/line/line-bot-sdk-go/linebot"

type Confirms struct {
	Title string
	Left  linebot.TemplateAction
	Right linebot.TemplateAction
}

func NewConfirms() Confirms {
	var c Confirms
	c.Title = "confirm"
	c.Left = &linebot.MessageTemplateAction{
		Label: "1",
		Text:  "Yes",
	}

	c.Right = &linebot.MessageTemplateAction{
		Label: "2",
		Text:  "No",
	}
	return c
}

func (c *Confirms) ConfirmsTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage("confirms template",
		&linebot.ConfirmTemplate{
			Text:    c.Title,
			Actions: []linebot.TemplateAction{c.Left, c.Right},
		})
}
