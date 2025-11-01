package ui

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

// PromptWithStatus shows a status message before prompting
func PromptWithStatus(status, label string) (string, error) {
	// Show status
	fmt.Printf("ℹ %s\n", status)

	prompt := promptui.Prompt{
		Label: label,
	}

	return prompt.Run()
}

// SelectWithStatus shows a status message before showing selection
func SelectWithStatus(status string, items []string, label string) (int, string, error) {
	// Show status
	fmt.Printf("ℹ %s\n", status)

	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	return prompt.Run()
}

// ShowStatus displays a simple status message
func ShowStatus(message string, success bool) {
	icon := "✗"
	if success {
		icon = "✓"
	}
	fmt.Fprintf(os.Stderr, "%s %s\n", icon, message)
}
