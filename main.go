package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/soeyusuke/bot/auth"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = auth.NewBot()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	http.HandleFunc("/callback", Callback)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func textHandler(message *linebot.TextMessage, replyToken string) error {
	var msg linebot.Message
	switch message.Text {
	case "buttons":

	case "confirm":
		t := linebot.NewConfirmTemplate(
			"Do it?",
			linebot.NewMessageTemplateAction("Yes", "Yes!"),
			linebot.NewMessageTemplateAction("No", "No!"),
		)
		msg = linebot.NewTemplateMessage("confirm template", t)

	case "carousel":
		t := linebot.NewCarouselTemplate(
			linebot.NewCarouselColumn(
				"https://avatars0.githubusercontent.com/u/24663217?s=400&v=4",
				"soeyu",
				"soeyusuke",
				linebot.NewURITemplateAction("github", "https://github.com/soeyusuke"),
				linebot.NewPostbackTemplateAction("Hi", "Nice!!", "", ""),
			),
			linebot.NewCarouselColumn(
				"https://avatars2.githubusercontent.com/u/29530489?s=400&v=4",
				"test",
				"test",
				linebot.NewPostbackTemplateAction("Noo!", "bad word!", "", ""),
				linebot.NewMessageTemplateAction("gg", "good game"),
			),
		)
		msg = linebot.NewTemplateMessage("carousel template", t)

	case "image carousel":
		imageURL := "https://tfsassets.azureedge.net/sampletry.jpg"
		t := linebot.NewImageCarouselTemplate(
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewURITemplateAction("Go to LINE", "https://line.me"),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", "", ""),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewMessageTemplateAction("Say message", "Rice=米"),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewDatetimePickerTemplateAction("datetime", "DATETIME", "datetime", "", "", ""),
			),
		)
		msg = linebot.NewTemplateMessage("image carousel template", t)
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
