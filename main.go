package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/soeyusuke/bot/auth"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = auth.NewBot(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
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
