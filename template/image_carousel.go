package template

import (
	"errors"

	"github.com/line/line-bot-sdk-go/linebot"
)

type ImageColumns struct {
	ImagePath string
	Action    linebot.TemplateAction
}

func NewImageColumns() ImageColumns {
	var col ImageColumns
	col.ImagePath = "https://tfsassets.azureedge.net/sampletry.jpg"
	col.Action = linebot.NewURITemplateAction("Go to LINE", "https://line.me")
	return col
}

type ImageCarousel struct {
	Columns []*linebot.ImageCarouselColumn
}

func NewImageCarousel(col []ImageColumns) (ImageCarousel, error) {
	if len(col) > 10 {
		return ImageCarousel{}, errors.New("columns must not be more than 10 items")
	}

	var c ImageCarousel
	for _, v := range col {
		c.Columns = append(c.Columns, &linebot.ImageCarouselColumn{
			ImageURL: v.ImagePath,
			Action:   v.Action,
		})
	}

	return c, nil
}

func (c ImageCarousel) CarouselTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage("image carousel",
		&linebot.ImageCarouselTemplate{
			Columns: c.Columns,
		})
}
