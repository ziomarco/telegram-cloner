package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageReceivedCallbackFunc = func(message *tgbotapi.Message)

func ListenToMessages(token string, cb MessageReceivedCallbackFunc, updateRange int) {
	client := initClient(token)
	config := tgbotapi.NewUpdate(updateRange)
	config.Timeout = 60
	updates := client.GetUpdatesChan(config)

	for update := range updates {
		if update.Message == nil {
			return
		}
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		cb(update.Message)
	}
}

func SendMessage(token string, to int64, msgContent string) {
	client := initClient(token)
	msg := tgbotapi.NewMessage(to, msgContent)
	client.Send(msg)
}

func initClient(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, "http://localhost:8081/user%s/%s")
	bot.Debug = false
	if err != nil {
		fmt.Println("Can't init TGBotClient")
		fmt.Println(err)
		return nil
	}
	return bot
}

func GetClient(token string) *tgbotapi.BotAPI {
	return initClient(token)
}
