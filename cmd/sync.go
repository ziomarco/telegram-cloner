/*
Copyright © 2023 brascioladisoia
*/
package cmd

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
	database "telegramcloner/db"
	"telegramcloner/http"
	"telegramcloner/prompt"
	"telegramcloner/telegram"
	"telegramcloner/textutils"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Start watching history and apply syncing rules",
	Long: `Use this command to start the syncing process by choosing
a source and a destination. Every message sent in source will be also
sent in destination.`,
	Run: sync,
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().Bool("debug", false, "Activate debug mode")
}

func sync(cmd *cobra.Command, args []string) {
	token := viper.GetString("telegram_user_token")
	if len(token) < 1 {
		log.Fatal("Please configure the CLI first!")
	}
	database.VerifyDatabase()
	chatList := http.GetChats(token)
	parsedOrigin, parsedDestination := prompt.PromptOriginDestination(chatList)
	stringToReplace, replacement := prompt.PromptSubstitution()
	charsToStrip := prompt.PromptCharsLengthToStrip()
	stripPhrases := prompt.PromptStripPhrase([]string{}, false)
	prompt.AskForSyncStartConfirmation(parsedOrigin, parsedDestination)

	log.Println(fmt.Sprintf("Using %s as origin and %s as destination", parsedOrigin, parsedDestination))
	if len(stringToReplace) > 0 {
		log.Println(fmt.Sprintf("Replacing \"%s\" with \"%s\"", stringToReplace, replacement))
	}

	var isDebugModeActive, _ = cmd.Flags().GetBool("debug")

	startSync(token, parsedOrigin, parsedDestination, stringToReplace, replacement, charsToStrip, stripPhrases, isDebugModeActive)
}

func startSync(token string, from string, to string, stringToReplace string, replacement string, charsToStrip int, stripPhrases []string, isDebugModeActive bool) {
	origin, _ := strconv.ParseInt(from, 10, 64)
	destination, _ := strconv.ParseInt(to, 10, 64)
	log.Println("Starting listening...")

	telegram.ListenToMessages(token, func(m *tgbotapi.Message) {
		handleMessage(token, m, origin, destination, stringToReplace, replacement, charsToStrip, stripPhrases, isDebugModeActive)
	}, 0)
}

func handleMessage(token string, m *tgbotapi.Message, origin int64, destination int64, stringToReplace string, replacement string, charsToStrip int, stripPhrases []string, isDebugModeActive bool) {
	log.Println("Handling message")

	if m.From == nil && m.SenderChat == nil {
		log.Println("Invalid update received!")
		return
	}
	if m.From != nil && m.From.ID != origin {
		return
	}
	if m.SenderChat != nil && m.SenderChat.ID != origin {
		return
	}

	var sender int64
	if m.From != nil {
		sender = m.From.ID
	}
	if m.SenderChat != nil {
		sender = m.SenderChat.ID
	}

	if isDebugModeActive {
		log.Printf("Debug mode active: saving a copy of message %d in database...", m.MessageID)
		debugMessageContent, _ := json.Marshal(m)
		out := database.RegisterHandledMessage(m.MessageID, string(debugMessageContent))
		if !out {
			log.Printf("Couldn't save a copy of original message %d in database", m.MessageID)
		}
	}

	// Message is not modified in no way, and it doesn't contain any linked message (replies)
	if len(stringToReplace) == 0 && charsToStrip == 0 && len(stripPhrases) == 0 && m.ReplyToMessage == nil {
		telegram.CopyMessage(token, origin, destination, *m)
		log.Println("Message forwarded")
		return
	}

	parsedMessage := textutils.ParseMessage(m, stringToReplace, replacement, charsToStrip, stripPhrases)
	var linkedMessageId int
	if m.ReplyToMessage != nil {
		linkedMessage := *database.FindMessage(m.ReplyToMessage.MessageID)
		if linkedMessage.ID < 0 {
			log.Println("Invalid message found in database")
		}
		linkedMessageId = linkedMessage.ForwardedMessageId
	}
	telegram.SendMessage(token, sender, destination, *parsedMessage, linkedMessageId)
	log.Println("Message forwarded")
}
