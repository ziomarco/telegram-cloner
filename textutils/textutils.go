package textutils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func ParseMessage(message *tgbotapi.Message, stringToReplace string, replacement string, charsToStrip int, stripPhrases []string) *tgbotapi.Message {
	if len(message.Caption) > 0 {
		message.Caption = parseText(message.Caption, stringToReplace, replacement, charsToStrip, stripPhrases)
	}

	if len(message.Text) > 0 {
		message.Text = parseText(message.Text, stringToReplace, replacement, charsToStrip, stripPhrases)
	}

	return message
}

func parseText(text string, stringToReplace string, replacement string, charsToStrip int, stripPhrases []string) string {
	text = strings.Replace(text, stringToReplace, replacement, -1)

	if charsToStrip > 0 {
		text = text[:len(text)-charsToStrip]
	}

	if len(stripPhrases) > 0 {
		for _, phraseToStrip := range stripPhrases {
			text = strings.Split(text, phraseToStrip)[0]
		}
	}

	return text
}
