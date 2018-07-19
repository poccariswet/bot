package template

import (
	"errors"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Carousel struct {
	Columns []*linebot.CarouselColumn
}

func NewColumns(btn []Buttons) (Carousel, error) {
	if len(btn) > 10 {
		return Carousel{}, errors.New("columns must not be more than 10 items")
	}

	var c Carousel
	for _, v := range btn {
		if len(v.Button) > 3 {
			return Carousel{}, errors.New("button actions must not be more than 3 items")
		}

		c.Columns = append(c.Columns, &linebot.CarouselColumn{
			ThumbnailImageURL: v.ImagePath,
			Title:             v.Title,
			Text:              v.SubTitle,
			Actions:           v.Button,
		})
	}
	return c, nil
}

func (c Carousel) CarouselTemplate() *linebot.TemplateMessage {
	return linebot.NewTemplateMessage("carousel",
		&linebot.CarouselTemplate{
			Columns: c.Columns,
		})
}
