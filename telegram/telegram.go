package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var globalClient tgbotapi.BotAPI

type MessageReceivedCallbackFunc = func(message *tgbotapi.Message)

func getClient(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, "http://localhost:8081/user%s/%s")
	bot.Debug = false
	if err != nil {
		log.Println("Can't init TGBotClient")
		log.Fatalln(err)
		return nil
	}
	return bot
}

func ListenToMessages(token string, cb MessageReceivedCallbackFunc, updateRange int) {
	client := getClient(token)

	config := tgbotapi.NewUpdate(updateRange)
	config.Timeout = 600

	updates := client.GetUpdatesChan(config)

	for update := range updates {
		if &update == nil {
			continue
		}
		if update.ChannelPost != nil {
			cb(update.ChannelPost)
		}
		if update.Message != nil {
			cb(update.Message)
		}
	}
}

func CopyMessage(token string, from int64, to int64, message tgbotapi.Message) {
	client := getClient(token)

	baseChat := tgbotapi.BaseChat{ChatID: to}
	_, err := client.CopyMessage(tgbotapi.CopyMessageConfig{
		BaseChat:            baseChat,
		FromChatID:          from,
		FromChannelUsername: "",
		MessageID:           message.MessageID,
		Caption:             message.Caption,
	})

	if err != nil {
		log.Panicln(err)
	}
}

func SendMessage(token string, from int64, to int64, message tgbotapi.Message) {
	client := getClient(token)

	var msgContent string

	if len(message.Text) > 0 {
		msgContent = message.Text
	}

	if len(message.Caption) > 0 {
		msgContent = message.Caption
	}

	if len(message.Photo) > 0 {
		photoId := message.Photo[len(message.Photo)-1].FileID
		photoMessageContent := tgbotapi.NewPhoto(to, tgbotapi.FileID(photoId))
		_, pErr := client.Send(photoMessageContent)
		if pErr != nil {
			log.Println("Error forwarding photo:")
			log.Println(pErr)
		}
	}

	msg := tgbotapi.NewMessage(to, msgContent)
	client.Send(msg)
}
