package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

type App struct {
	bot *linebot.Client
}

func main() {
	app, err := NewApp()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	http.HandleFunc("/callback", app.Callback)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func NewApp() (*App, error) {
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		return nil, err
	}

	return &App{bot: bot}, nil
}

func (app *App) Callback(w http.ResponseWriter, req *http.Request) {
	events, err := app.bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		log.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := app.textHandler(message, event.ReplyToken); err != nil {
					log.Println(err)
				}
			}

		case linebot.EventTypePostback:
			data := event.Postback.Data
			if data == "DATE" || data == "TIME" || data == "DATETIME" {
				data += fmt.Sprintf("(%v)", *event.Postback.Params)
			}
			if _, err = app.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(data)).Do(); err != nil {
				log.Println(err)
			}
		}
	}
}

func (app *App) textHandler(message *linebot.TextMessage, replyToken string) error {
	var msg linebot.Message
	switch message.Text {
	case "buttons":
		t := linebot.NewButtonsTemplate(
			"",                                                                    // image path
			"My button sample",                                                    // ButtonsTemplate Title
			"Hello, my button",                                                    // ButtonsTemplate SubTitle
			linebot.NewPostbackTemplateAction("1", "push button1", "", "button1"), // (label, data, text, displayText)
			linebot.NewPostbackTemplateAction("2", "push button2", "", "button2"), // button template can't use text and
			linebot.NewPostbackTemplateAction("3", "push button3", "", ""),        //displayText don't use at same time
		)
		msg = linebot.NewTemplateMessage("buttons template", t)

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
		if _, err := app.bot.ReplyMessage(replyToken, msg).Do(); err != nil {
			return err
		}
	}
	return nil
}
