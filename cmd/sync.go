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
	"golang.org/x/exp/slices"
	"log"
	"os"
	"telegramcloner/config"
	database "telegramcloner/db"
	"telegramcloner/http"
	"telegramcloner/telegram"
	"telegramcloner/textutils"
	"time"
)

const KDetectionInterval = 20

var logPath, _ = os.Getwd()
var KLogPath = logPath + "/telegramcloner" + time.Now().Format("2006-01-02--15-04") + ".log"

var outfile, _ = os.Create(KLogPath)
var logger = log.New(outfile, "[syncCmd] ", log.LstdFlags|log.Lshortfile)

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
	viper.WatchConfig()
}

func sync(cmd *cobra.Command, args []string) {
	token := viper.GetString("telegram_user_token")
	if len(token) < 1 {
		logger.Fatal("Please login into telegram using \"configure\" command first")
	}
	// database.VerifyDatabase()

	var isDebugModeActive, _ = cmd.Flags().GetBool("debug")

	config.CheckConfigFileExistence()
	stringToReplace := viper.GetString(KStringToReplace)
	replacement := viper.GetString(KReplacement)
	charsToStrip := viper.GetInt(KCharsToStrip)
	stripPhrases := viper.GetStringSlice(KStripPhrases)

	startSync(token, stringToReplace, replacement, charsToStrip, stripPhrases, isDebugModeActive)
}

func detectIdChanges(isDebugModeActive bool) {
	ticker := time.NewTicker(KDetectionInterval * time.Second)
	for _ = range ticker.C {
		if isDebugModeActive {
			logger.Println("[detectIdChanges] Checking if source and destination id changed")
		}
		var origin http.ChatListItem
		var destination http.ChatListItem
		viper.UnmarshalKey(KOrigin, &origin)
		viper.UnmarshalKey(KDestination, &destination)
		updatedChats := http.GetChats(viper.GetString(KTelegramUserToken))

		originIdx := slices.IndexFunc(updatedChats, func(el http.ChatListItem) bool {
			return el.Name == origin.Name
		})
		destinationIdx := slices.IndexFunc(updatedChats, func(el http.ChatListItem) bool {
			return el.Name == destination.Name
		})

		if updatedChats[originIdx].Id != origin.Id {
			if isDebugModeActive {
				logger.Println("[detectIdChanges] Origin ID changed, updating configuration...")
			}
			viper.Set(KOrigin, updatedChats[originIdx])
			viper.Set(KParsedOrigin, updatedChats[originIdx].Id)
		} else {
			if isDebugModeActive {
				logger.Println("[detectIdChanges] Origin ID unchanged, skipping...")
			}
		}

		if updatedChats[destinationIdx].Id != destination.Id {
			if isDebugModeActive {
				logger.Println("[detectIdChanges] Destination ID changed, updating configuration...")
			}
			viper.Set(KDestination, updatedChats[destinationIdx])
			viper.Set(KParsedDestination, updatedChats[destinationIdx].Id)
		} else {
			if isDebugModeActive {
				logger.Println("[detectIdChanges] Destination ID unchanged, skipping...")
			}
		}
	}
}

func startSync(token string, stringToReplace string, replacement string, charsToStrip int, stripPhrases []string, isDebugModeActive bool) {
	logger.Println("Starting listening...")
	logger.Printf("Using %s as origin and %s as destination", viper.Get(KOrigin), viper.Get(KDestination))

	go detectIdChanges(isDebugModeActive)

	telegram.ListenToMessages(token, func(m *tgbotapi.Message, isUpdate bool) {
		origin := int64(viper.GetInt(KParsedOrigin))
		destination := int64(viper.GetInt(KParsedDestination))
		handleMessage(token, m, isUpdate, origin, destination, stringToReplace, replacement, charsToStrip, stripPhrases, isDebugModeActive)
	}, 0)
}

func handleMessage(token string, m *tgbotapi.Message, isUpdate bool, origin int64, destination int64, stringToReplace string, replacement string, charsToStrip int, stripPhrases []string, isDebugModeActive bool) {
	logger.Println("Handling message")

	if m.From == nil && m.SenderChat == nil {
		logger.Println("Invalid update received!")
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
		data, _ := json.Marshal(m)
		fmt.Println("Actively handling message")
		fmt.Println(string(data))
		logger.Printf("Debug mode active: saving a copy of message %d in database...", m.MessageID)
		debugMessageContent, _ := json.Marshal(m)
		out := database.RegisterHandledMessage(m.MessageID, string(debugMessageContent))
		if !out {
			logger.Printf("Couldn't save a copy of original message %d in database", m.MessageID)
		}
	}

	parsedMessage := textutils.ParseMessage(m, stringToReplace, replacement, charsToStrip, stripPhrases)

	if isUpdate {
		fmt.Println("Handling an update...")
		dbMessage := database.FindMessage(m.MessageID)
		if dbMessage.ID == 0 {
			logger.Printf("Can't edit message %d because no match was found in DB", m.MessageID)
			return
		}
		forwardedChatId := dbMessage.DestinationChatId
		forwardedMessageId := dbMessage.ForwardedMessageId
		telegram.EditMessage(token, int64(forwardedChatId), forwardedMessageId, textutils.ExtractTextFromMessage(parsedMessage))
		return
	}

	var linkedMessageId int
	if m.ReplyToMessage != nil {
		linkedMessage := *database.FindMessage(m.ReplyToMessage.MessageID)
		if linkedMessage.ID == 0 {
			logger.Println("Invalid message found in database")
			logger.Printf("forwarding anyway message %d \n", m.MessageID)
		}
		linkedMessageId = linkedMessage.ForwardedMessageId
	}
	telegram.SendMessage(token, sender, destination, *parsedMessage, linkedMessageId)
	logger.Println("Message forwarded")
}
