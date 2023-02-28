/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"regexp"
	"strings"
	"telegramcloner/http"
	"telegramcloner/prompt"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	confirm := *prompt.Select(prompt.SelectPromptContent{
		ErrorMsg: "Invalid response!",
		Label:    fmt.Sprintf("Do you want to enable sync from %s to %s? (yes/no)", parsedOrigin, parsedDestination),
		Items:    []string{"yes", "no"},
	})

	if confirm != "yes" {
		log.Println("Ok, interrupting...")
		os.Exit(1)
	}

	log.Println(fmt.Sprintf("Using %s as origin and %s as destination", parsedOrigin, parsedDestination))
}
