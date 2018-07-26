package main

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/soeyusuke/bot/template"
)

func textHandler(message *linebot.TextMessage, replyToken string) error {
	var msg linebot.Message
	switch message.Text {
	case "buttons":
		btn := template.NewButtons()
		msg = btn.ButtonsTemplate()

	case "confirm":
		confirm := template.NewConfirms()
		msg = confirm.ConfirmsTemplate()

	case "carousel":
		var btns []template.Buttons
		btns = append(btns, template.NewButtons())
		btns = append(btns, template.NewButtons())
		carousel, err := template.NewColumns(btns)
		if err != nil {
			return err
		}
		msg = carousel.CarouselTemplate()

	case "image carousel":
		var col []template.ImageColumns
		col = append(col, template.NewImageColumns())
		col = append(col, template.NewImageColumns())
		col = append(col, template.NewImageColumns())

		c, err := template.NewImageCarousel(col)
		if err != nil {
			return err
		}
		msg = c.CarouselTemplate()

	case "flex json":
		//TODO: adding flex_json template func

		contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
		if err != nil {
			return err
		}
		if _, err := bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("Flex message alt text", contents),
		).Do(); err != nil {
			return err
		}

	default:
		log.Println(message.Text)
	}
	if msg != nil {
		if _, err := bot.ReplyMessage(replyToken, msg).Do(); err != nil {
			return err
		}
	}
	return nil
}
