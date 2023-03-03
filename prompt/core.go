package prompt

import (
	"errors"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"regexp"
)

type InputPromptContent struct {
	ErrorMsg      string
	Label         string
	ValidateFunc  func(string) error
	ValidateRegex string
}

type SelectPromptContent struct {
	ErrorMsg string
	Label    string
	Items    []string
}

func Input(pc InputPromptContent) string {
	validate := func(input string) error {
		if pc.ValidateFunc != nil {
			return pc.ValidateFunc(input)
		}
		if len(pc.ValidateRegex) > 1 {
			match, _ := regexp.MatchString(pc.ValidateRegex, input)
			if !match {
				return errors.New(pc.ErrorMsg)
			}
		}
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func Select(pc SelectPromptContent) *string {
	prompt := promptui.Select{
		Label: pc.Label,
		Items: pc.Items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return nil
	}

	return &result
}
