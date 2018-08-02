package template

import (
	"errors"

	"github.com/line/line-bot-sdk-go/linebot"
)

type ImageColumns struct {
	ImagePath string
	Action    linebot.TemplateAction
}

// new defaul image columns template
func NewImageColumns() *ImageColumns {
	return &ImageColumns{
		ImagePath: "https://cdn2.iconfinder.com/data/icons/file-8/128/file_3-512.png",
		Action:    nil,
	}
}

// set Image columns
func (c *ImageColumns) SetImageAction(action linebot.TemplateAction) {
	c.Action = action
}

func (c *ImageColumns) SetImagePath(imagepath string) {
	c.ImagePath = imagepath
}

type ImageCarousel struct {
	Columns []*linebot.ImageCarouselColumn
}

func NewImageCarousel() *ImageCarousel {
	return &ImageCarousel{}
}

// set image carousel
func (c *ImageCarousel) SetImageCarousel(col ...*ImageColumns) error {
	if len(col) > 10 {
		return errors.New("columns must not be more than 10 items")
	}

	for _, v := range col {
		c.Columns = append(c.Columns, &linebot.ImageCarouselColumn{
			ImageURL: v.ImagePath,
			Action:   v.Action,
		})
	}

	return nil
}

func (c ImageCarousel) CarouselTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage("image carousel",
		&linebot.ImageCarouselTemplate{
			Columns: c.Columns,
		})
}
