package template

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

type Buttons struct {
	ImagePath string
	Title     string
	SubTitle  string
	Button    []linebot.TemplateAction
}

//func NewButtons() Buttons {
//	var btn Buttons
//	btn.ImagePath = "https://avatars1.githubusercontent.com/u/24663217?s=460&v=4"
//	btn.Title = "buttons"
//	btn.SubTitle = "sub_buttons"
//	btn.Button = append(btn.Button, linebot.NewURITemplateAction("Go to line.me", "https://line.me"))
//	btn.Button = append(btn.Button, linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", "", "hello こんにちは"))
//	btn.Button = append(btn.Button, linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは", ""))
//
//	return btn
//}

//default fill
func NewButtons() Buttons {
	return Buttons{
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

//TODO: Append the some templatate action to Buttons Button
func (btn *Buttons) AddButtons(button linebot.TemplateAction) {
	btn.Button = append(btn.Button, button)
}

func (btn *Buttons) SetImagePath(imagepath string) {
	btn.ImagePath = imagepath
}

func (btn *Buttons) SetTitle(title string) {
	btn.Title = title
}

func (btn *Buttons) SetSubTitle(subtitle string) {
	btn.SubTitle = subtitle
}

func (btn *Buttons) ButtonsTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage("buttons template",
		&linebot.ButtonsTemplate{
			ThumbnailImageURL: btn.ImagePath,
			Title:             btn.Title,
			Text:              btn.SubTitle,
			Actions:           btn.Button,
		})
}
