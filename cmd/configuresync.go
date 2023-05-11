package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"telegramcloner/config"
	"telegramcloner/http"
	"telegramcloner/prompt"
)

var configureSyncCmd = &cobra.Command{
	Use:   "configure-sync",
	Short: "Launch the configuration process for saving syncing rules",
	Long: `Use this command to configure the syncing progress by choosing
a source, a destination and processing rules. It's needed for saving the
syncing configurations into a file that "sync" command will read.'`,
	Run: configureSync,
}

func init() {
	rootCmd.AddCommand(configureSyncCmd)
}

const (
	KOrigin            = "raw_origin"
	KDestination       = "raw_destination"
	KParsedOrigin      = "parsed_origin"
	KParsedDestination = "parsed_destination"
	KStringToReplace   = "string_to_replace"
	KReplacement       = "replacement"
	KCharsToStrip      = "chars_to_strip"
	KStripPhrases      = "strip_phrases"
)

func configureSync(cmd *cobra.Command, args []string) {
	token := viper.GetString("telegram_user_token")
	if len(token) < 1 {
		log.Fatal("Please login into telegram using \"configure\" command first")
	}
	chatList := http.GetChats(token)
	origin, destination := prompt.PromptOriginDestination(chatList)
	stringToReplace, replacement := prompt.PromptSubstitution()
	charsToStrip := prompt.PromptCharsLengthToStrip()
	stripPhrases := prompt.PromptStripPhrase([]string{}, false)

	log.Println(fmt.Sprintf("Using %s as origin and %s as destination", origin.Name, destination.Name))
	if len(stringToReplace) > 0 {
		log.Println(fmt.Sprintf("Replacing \"%s\" with \"%s\"", stringToReplace, replacement))
	}

	viper.Set(KOrigin, origin)
	viper.Set(KDestination, destination)
	viper.Set(KParsedOrigin, origin.Id)
	viper.Set(KParsedDestination, destination.Id)
	viper.Set(KStringToReplace, stringToReplace)
	viper.Set(KReplacement, replacement)
	viper.Set(KCharsToStrip, charsToStrip)
	viper.Set(KStripPhrases, stripPhrases)
	config.WriteConfig(viper.AllSettings())

	log.Println("Config saved!")
}
