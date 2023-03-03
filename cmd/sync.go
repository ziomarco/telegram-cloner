/*
Copyright © 2023 brascioladisoia
*/
package cmd

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"telegramcloner/http"
	"telegramcloner/prompt"
	"telegramcloner/telegram"
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
}

func sync(cmd *cobra.Command, args []string) {
	token := viper.GetString("telegram_user_token")
	if len(token) < 1 {
		log.Fatal("Please configure the CLI first!")
	}
	chatList := http.GetChats(token)
	parsedOrigin, parsedDestination := promptOriginDestination(chatList)
	stringToReplace, replacement := promptSubstitution()
	askForConfirm(parsedOrigin, parsedDestination)

	log.Println(fmt.Sprintf("Using %s as origin and %s as destination", parsedOrigin, parsedDestination))
	if len(stringToReplace) > 0 {
		log.Println(fmt.Sprintf("Replacing \"%s\" with \"%s\"", stringToReplace, replacement))
	}
	startSync(token, parsedOrigin, parsedDestination, stringToReplace, replacement)
}

func startSync(token string, from string, to string, stringToReplace string, replacement string) {
	origin, _ := strconv.ParseInt(from, 10, 64)
	destination, _ := strconv.ParseInt(to, 10, 64)
	log.Println("Starting listening...")
	telegram.ListenToMessages(token, func(m *tgbotapi.Message) {
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
		if len(stringToReplace) == 0 {
			telegram.CopyMessage(token, origin, destination, *m)
			log.Println("Message forwarded")
			return
		}

		parsedMessage := parseMessage(m, stringToReplace, replacement)
		telegram.SendMessage(token, sender, destination, *parsedMessage)
		log.Println("Message forwarded")
	}, 0)
}

func parseMessage(message *tgbotapi.Message, stringToReplace string, replacement string) *tgbotapi.Message {
	if len(message.Caption) > 0 {
		message.Caption = strings.Replace(message.Caption, stringToReplace, replacement, -1)
		return message
	}
	if len(message.Text) > 0 {
		message.Text = strings.Replace(message.Text, stringToReplace, replacement, -1)
		return message
	}
	return message
}

func promptOriginDestination(chatList []string) (string, string) {
	origin := *prompt.Select(prompt.SelectPromptContent{
		ErrorMsg: "No origin selected!",
		Label:    "Select an origin",
		Items:    chatList,
	})
	destination := *prompt.Select(prompt.SelectPromptContent{
		ErrorMsg: "No origin selected!",
		Label:    "Select a destination",
		Items:    chatList,
	})
	r := regexp.MustCompile("\\[([^\\][]*)]")
	parsedOrigin := r.FindString(origin)
	parsedOrigin = strings.Trim(parsedOrigin, "[")
	parsedOrigin = strings.Trim(parsedOrigin, "]")
	parsedDestination := r.FindString(destination)
	parsedDestination = strings.Trim(parsedDestination, "[")
	parsedDestination = strings.Trim(parsedDestination, "]")
	return parsedOrigin, parsedDestination
}

func promptSubstitution() (string, string) {
	confirm := *prompt.Select(prompt.SelectPromptContent{
		ErrorMsg: "Invalid input!",
		Label:    fmt.Sprintf("Do you want to substitute some string in message forwarding?"),
		Items:    []string{"yes", "no"},
	})

	if confirm != "yes" {
		return "", ""
	}

	textToReplace := prompt.Input(prompt.InputPromptContent{
		ErrorMsg:      "Invalid input!",
		Label:         "Insert text to replace text",
		ValidateFunc:  nil,
		ValidateRegex: "",
	})

	replacement := prompt.Input(prompt.InputPromptContent{
		ErrorMsg: "Invalid input!",
		Label:    "Insert replacement",
		ValidateFunc: func(s string) error {
			return nil
		},
		ValidateRegex: "",
	})

	return textToReplace, replacement
}

func askForConfirm(origin string, destination string) {
	confirm := *prompt.Select(prompt.SelectPromptContent{
		ErrorMsg: "Invalid response!",
		Label:    fmt.Sprintf("Do you want to enable sync from %s to %s? (yes/no)", origin, destination),
		Items:    []string{"yes", "no"},
	})

	if confirm != "yes" {
		log.Println("Ok, interrupting...")
		os.Exit(1)
	}
}
