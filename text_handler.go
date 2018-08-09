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
		if err := btn.AddButtons(
			linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", "hello こんにちは"),
			linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
			linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
			linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
		); err != nil {
			return err
		}
		msg = btn.ButtonsTemplate()

	case "confirm":
		confirm := template.NewConfirms()
		msg = confirm.ConfirmsTemplate()

	case "carousel":
		carousel := template.NewCarousel()
		btn := template.NewButtons()
		if err := btn.AddButtons(
			linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", "hello こんにちは"),
			linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
			linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
		); err != nil {
			return err
		}

		if err := carousel.SetColumns(btn, btn); err != nil {
			return err
		}
		log.Println(carousel.Columns)
		msg = carousel.CarouselTemplate()

	case "image carousel":
		col := template.NewImageColumns()
		col.SetImageAction(linebot.NewURIAction("Go to LINE", "https://line.me"))

		c := template.NewImageCarousel()
		if err := c.SetImageCarousel(col, col, col); err != nil {
			return err
		}
		msg = c.CarouselTemplate()

	case "flex json":
		f := template.NewFlexMessage()
		f.SetHero(template.ImageComponent{
			Type:        "image",
			URL:         "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
			Size:        "full",
			AspectRatio: "20:13",
			AspectMode:  "cover",
			Action: linebot.URIAction{
				URI: "http://linecorp.com/",
			},
		})

		f.SetBody(template.BoxComponent{
			Type:   "box",
			Layout: "vertical",
			Contents: []template.FlexComponent{
				template.TextComponent{
					Type:   "text",
					Text:   "Brown Cafe",
					Weight: "bold",
					Size:   "xl",
				},
				template.BoxComponent{
					Type:   "box",
					Layout: "baseline",
					Margin: "md",
					Contents: []template.FlexComponent{
						template.IconComponent{
							Type: "icon",
							Size: "sm",
							URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png",
						},
						template.IconComponent{
							Type: "icon",
							Size: "sm",
							URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png",
						},
						template.IconComponent{
							Type: "icon",
							Size: "sm",
							URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png",
						},
						template.IconComponent{
							Type: "icon",
							Size: "sm",
							URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png",
						},
						template.IconComponent{
							Type: "icon",
							Size: "sm",
							URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png",
						},
						template.TextComponent{
							Type:   "text",
							Text:   "4.0",
							Size:   "sm",
							Color:  "#9999999",
							Margin: "md",
							Flex:   0,
						},
					},
				},
				template.BoxComponent{
					Type:    "box",
					Layout:  "vertical",
					Margin:  "lg",
					Spacing: "sm",
					Contents: []template.FlexComponent{
						template.BoxComponent{
							Type:    "box",
							Layout:  "baseline",
							Spacing: "sm",
							Contents: []template.FlexComponent{
								template.TextComponent{
									Type:  "text",
									Text:  "Place",
									Color: "#aaaaaa",
									Size:  "sm",
									Flex:  1,
								},
								template.TextComponent{
									Type:  "text",
									Text:  "demo demo demo",
									Wrap:  true,
									Color: "#666666",
									Size:  "sm",
									Flex:  5,
								},
							},
						},
						template.BoxComponent{
							Type:    "box",
							Layout:  "baseline",
							Spacing: "sm",
							Contents: []template.FlexComponent{
								template.TextComponent{
									Type:  "text",
									Text:  "Time",
									Color: "#aaaaaa",
									Size:  "sm",
									Flex:  1,
								},
								template.TextComponent{
									Type:  "text",
									Text:  "10:00 - 23:00",
									Wrap:  true,
									Color: "#666666",
									Size:  "sm",
									Flex:  5,
								},
							},
						},
					},
				},
			},
		})

		f.SetFooter(template.BoxComponent{
			Type:    "box",
			Layout:  "vertical",
			Spacing: "sm",
			Contents: []template.FlexComponent{
				template.ButtonComponent{
					Type:   "button",
					Style:  "link",
					Height: "sm",
					Action: linebot.URIAction{
						Label: "CALL",
						URI:   "https://linecorp.com",
					},
				},
				template.ButtonComponent{
					Type:   "button",
					Style:  "link",
					Height: "sm",
					Action: linebot.URIAction{
						Label: "WEBSITE",
						URI:   "https://linecorp.com",
					},
				},
			},
			Flex: 0,
		})

		msg, err = f.FlexTemplate()
		if err != nil {
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
