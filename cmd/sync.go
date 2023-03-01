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
	askForConfirm(parsedOrigin, parsedDestination)

	log.Println(fmt.Sprintf("Using %s as origin and %s as destination", parsedOrigin, parsedDestination))
	startSync(token, parsedOrigin, parsedDestination)
}

func startSync(token string, from string, to string) {
	origin, _ := strconv.ParseInt(from, 10, 64)
	destination, _ := strconv.ParseInt(to, 10, 64)
	log.Println("Starting listening...")
	telegram.ListenToMessages(token, func(m *tgbotapi.Message) {
		log.Println("Handling message")
		if m.From.ID != origin {
			return
		}
		telegram.SendMessage(token, destination, m.Text)
		log.Println("Message forwarded")
	}, 0)
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
