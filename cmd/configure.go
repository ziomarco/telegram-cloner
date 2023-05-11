/*
Copyright © 2023 brascioladisoia
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"telegramcloner/config"
	"telegramcloner/http"
	"telegramcloner/prompt"
)

const KTelegramUserToken = "telegram_user_token"

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Login as user on Telegram Core",
	Long:  `Nothing to add`,
	Run:   configure,
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

func configure(cmd *cobra.Command, args []string) {
	phoneNumber := prompt.Input(prompt.InputPromptContent{
		ErrorMsg:      "Insert a valid phone number",
		Label:         "Phone Number (with prefix)",
		ValidateFunc:  nil,
		ValidateRegex: "\\+(9[976]\\d|8[987530]\\d|6[987]\\d|5[90]\\d|42\\d|3[875]\\d|\n2[98654321]\\d|9[8543210]|8[6421]|6[6543210]|5[87654321]|\n4[987654310]|3[9643210]|2[70]|7|1)\\d{1,14}$",
	})
	loginResponse := http.Login(phoneNumber)
	if loginResponse == nil {
		log.Fatal("Impossible to login")
	}
	otp := prompt.Input(prompt.InputPromptContent{
		ErrorMsg:      "Insert correct otp",
		Label:         "Insert the OTP Telegram sent to you",
		ValidateFunc:  nil,
		ValidateRegex: "",
	})
	confirmLoginResponse := http.ConfirmLogin(loginResponse.Result.Token, otp)
	if confirmLoginResponse == nil || confirmLoginResponse.Result.AuthorizationState != "ready" {
		log.Fatal("Impossible to confirm login")
	}
	viper.Set(KTelegramUserToken, loginResponse.Result.Token)
	err := config.WriteConfig(viper.AllSettings())
	if err != nil {
		log.Fatalln("Impossible to write configuration!")
	}
	log.Println("Login succeeded!")
}
