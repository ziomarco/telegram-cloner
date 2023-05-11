package prompt

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"strconv"
	"strings"
	"telegramcloner/http"
)

func PromptOriginDestination(chatList []http.ChatListItem) (http.ChatListItem, http.ChatListItem) {
	var parsedChatList []string
	for _, chat := range chatList {
		parsedChatList = append(parsedChatList, fmt.Sprintf("%s△%s△%s", chat.Name, chat.Type, strconv.Itoa(chat.Id)))
	}

	strOrigin := *Select(SelectPromptContent{
		ErrorMsg: "No origin selected!",
		Label:    "Select an origin",
		Items:    parsedChatList,
	})

	strDestination := *Select(SelectPromptContent{
		ErrorMsg: "No origin selected!",
		Label:    "Select a destination",
		Items:    parsedChatList,
	})

	originIdx := slices.IndexFunc(chatList, func(el http.ChatListItem) bool {
		split := strings.Split(strOrigin, "△")
		parsedId, _ := strconv.Atoi(split[2])
		return split[0] == el.Name && split[1] == el.Type && parsedId == el.Id
	})

	destinationIdx := slices.IndexFunc(chatList, func(el http.ChatListItem) bool {
		split := strings.Split(strDestination, "△")
		parsedId, _ := strconv.Atoi(split[2])
		return split[0] == el.Name && split[1] == el.Type && parsedId == el.Id
	})

	return chatList[originIdx], chatList[destinationIdx]
}

func PromptSubstitution() (string, string) {
	confirm := *Select(SelectPromptContent{
		ErrorMsg: "Invalid input!",
		Label:    fmt.Sprintf("Do you want to substitute some string in message forwarding?"),
		Items:    []string{"yes", "no"},
	})

	if confirm != "yes" {
		return "", ""
	}

	textToReplace := Input(InputPromptContent{
		ErrorMsg:      "Invalid input!",
		Label:         "Insert text to replace text",
		ValidateFunc:  nil,
		ValidateRegex: "",
	})

	replacement := Input(InputPromptContent{
		ErrorMsg: "Invalid input!",
		Label:    "Insert replacement",
		ValidateFunc: func(s string) error {
			return nil
		},
		ValidateRegex: "",
	})

	return textToReplace, replacement
}
func PromptCharsLengthToStrip() int {
	confirm := *Select(SelectPromptContent{
		ErrorMsg: "Invalid input!",
		Label:    fmt.Sprintf("Do you want to cut out some characters in message forwarding?"),
		Items:    []string{"yes", "no"},
	})

	if confirm != "yes" {
		return 0
	}

	charsToStripString := Input(InputPromptContent{
		ErrorMsg: "Invalid input!",
		Label:    "Insert number of chars to strip (from the end)",
		ValidateFunc: func(s string) error {
			val, _ := strconv.Atoi(s)
			if !(val > 0) {
				return errors.New("invalid number")
			}
			return nil
		},
		ValidateRegex: "",
	})

	charsToStrip, _ := strconv.Atoi(charsToStripString)

	return charsToStrip
}

func PromptStripPhrase(a []string, isFunctionInvokedAsRecursive bool) []string {
	if !isFunctionInvokedAsRecursive {
		confirm := *Select(SelectPromptContent{
			ErrorMsg: "Invalid input!",
			Label:    fmt.Sprintf("Do you want to strip out some phrases in message forwarding?"),
			Items:    []string{"yes", "no"},
		})

		if confirm != "yes" {
			return a
		}
	}

	stripPhrase := Input(InputPromptContent{
		ErrorMsg: "Invalid input!",
		Label:    "Insert phrase to strip from (first part will be taken)",
		ValidateFunc: func(s string) error {
			if len(s) == 0 {
				return errors.New("invalid phrase to strip")
			}
			return nil
		},
		ValidateRegex: "",
	})

	a = append(a, stripPhrase)

	confirmAnother := *Select(SelectPromptContent{
		ErrorMsg: "Invalid input!",
		Label:    fmt.Sprintf("Do you want to add another phrase to strip?"),
		Items:    []string{"yes", "no"},
	})

	if confirmAnother == "yes" {
		return PromptStripPhrase(a, true)
	}

	return a
}

func AskForSyncStartConfirmation(origin string, destination string) {
	confirm := *Select(SelectPromptContent{
		ErrorMsg: "Invalid response!",
		Label:    fmt.Sprintf("Do you want to enable sync from %s to %s? (yes/no)", origin, destination),
		Items:    []string{"yes", "no"},
	})

	if confirm != "yes" {
		log.Println("Ok, interrupting...")
		os.Exit(1)
	}
}
