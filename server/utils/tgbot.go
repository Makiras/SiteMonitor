package utils

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	tgbot *tgbotapi.BotAPI
)

// SendTgNotifaction provide send notifaction to telegram users
func SendTgNotifaction(to int64, msg string) error {
	tgmsg := tgbotapi.NewMessage(to, msg)
	_, err := tgbot.Send(tgmsg)
	return err
}

// InitTgBotHandler provides initial for tgbot client handler
func InitTgBotHandler() error {
	var err error
	tgbot, err = tgbotapi.NewBotAPI("") //todo: config for tg bot
	if err != nil {
		log.Print(err)
	}
	return err
}

// todo: handle tg user requests
