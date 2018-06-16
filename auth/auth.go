package auth

import (
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func NewBot() (*linebot.Client, error) {
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		return nil, err
	}

	return bot, nil
}
