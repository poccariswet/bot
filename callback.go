package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func postbackHandle(data string) error {
	return nil
}

func Callback(w http.ResponseWriter, req *http.Request) {
	events, err := bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
	}

	for _, event := range events {
		//log of event
		log.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := textHandler(message, event.ReplyToken); err != nil {
					log.Println(err)
				}
			}

		case linebot.EventTypePostback:
			data := event.Postback.Data
			if err := postbackHandle(data); err != nil {
				log.Println(err)
			}
		}
	}
}
