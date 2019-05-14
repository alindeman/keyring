package keyring

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

type PromptFunc func(string) (string, error)

func TerminalPrompt(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	b, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(b), nil
}

func FixedStringPrompt(value string) PromptFunc {
	return func(_ string) (string, error) {
		return value, nil
	}
}
