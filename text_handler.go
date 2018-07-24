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

	case "flex":
		// {
		//   "type": "bubble",
		//   "body": {
		//     "type": "box",
		//     "layout": "horizontal",
		//     "contents": [
		//       {
		//         "type": "text",
		//         "text": "Hello,"
		//       },
		//       {
		//         "type": "text",
		//         "text": "World!"
		//       }
		//     ]
		//   }
		// }
		contents := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Body: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "Hello,",
					},
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "World!",
					},
				},
			},
		}
		if _, err := bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("Flex message alt text", contents),
		).Do(); err != nil {
			return err
		}

	case "flex carousel":
		// {
		//   "type": "carousel",
		//   "contents": [
		//     {
		//       "type": "bubble",
		//       "body": {
		//         "type": "box",
		//         "layout": "vertical",
		//         "contents": [
		//           {
		//             "type": "text",
		//             "text": "First bubble"
		//           }
		//         ]
		//       }
		//     },
		//     {
		//       "type": "bubble",
		//       "body": {
		//         "type": "box",
		//         "layout": "vertical",
		//         "contents": [
		//           {
		//             "type": "text",
		//             "text": "Second bubble"
		//           }
		//         ]
		//       }
		//     }
		//   ]
		// }
		contents := &linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
			Contents: []*linebot.BubbleContainer{
				&linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "First bubble",
							},
						},
					},
				},
				&linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "Second bubble",
							},
						},
					},
				},
			},
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
