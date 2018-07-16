package auth

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func NewBot(secret, token string) (*linebot.Client, error) {
	bot, err := linebot.New(
		secret,
		token,
	)
	if err != nil {
		return nil, err
	}

	return bot, nil
}
