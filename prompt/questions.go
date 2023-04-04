package prompt

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PromptOriginDestination(chatList []string) (string, string) {
	origin := *Select(SelectPromptContent{
		ErrorMsg: "No origin selected!",
		Label:    "Select an origin",
		Items:    chatList,
	})

	destination := *Select(SelectPromptContent{
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
