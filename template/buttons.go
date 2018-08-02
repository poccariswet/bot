package template

import (
	"errors"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Buttons struct {
	ImagePath string
	Title     string
	SubTitle  string
	Button    []linebot.TemplateAction
}

// new default buttons template
func NewButtons() *Buttons {
	return &Buttons{
		ImagePath: "https://cdn2.iconfinder.com/data/icons/file-8/128/file_3-512.png",
		Title:     "title",
		SubTitle:  "sub title",
		Button:    nil,
	}
}

//TODO: Read to some type's file, so put the value in btn (Buttons)
func (btn *Buttons) Read() error {
	return nil
}

// Append the some templatate action to Buttons Button
func (btn *Buttons) AddButtons(buttons ...linebot.TemplateAction) error {
	if len(btn.Button) >= 4 {
		return errors.New("must not be more than 4 items")
	}

	for _, v := range buttons {
		btn.Button = append(btn.Button, v)
	}
	return nil
}

// set the buttons params
func (btn *Buttons) SetButtons(imagepath, title, subtitle string, buttons ...linebot.TemplateAction) error {
	btn.ImagePath = imagepath
	btn.Title = title
	btn.SubTitle = subtitle

	if len(buttons) > 4 {
		return errors.New("must not be more than 4 items")
	}

	btn.Button = nil
	for _, v := range buttons {
		btn.Button = append(btn.Button, v)
	}

	return nil
}

// implement button template to template message is for sending line
func (btn *Buttons) ButtonsTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage("buttons template",
		&linebot.ButtonsTemplate{
			ThumbnailImageURL: btn.ImagePath,
			Title:             btn.Title,
			Text:              btn.SubTitle,
			Actions:           btn.Button,
		})
}
