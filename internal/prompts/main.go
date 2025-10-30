package selects

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

func SelectProgrammingLanguage() (string, error) {
	programmingLanguages := []string{"Javascript", "Go"}

	prompt := promptui.Select{
		Label: "Select the programming language",
		Items: programmingLanguages,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("prompt failed %v", err)
	}

	return result, nil
}

func SelectJsFramework() (string, error) {
	frameworks := []string{"React", "Vue", "Angular"}

	prompt := promptui.Select{
		Label: "Select the Javascript framework",
		Items: frameworks,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("prompt failed %v", err)
	}

	return result, nil
}

func SelectGoFramework() (string, error) {
	goFrameworks := []string{"Gin", "Echo", "Fiber"}

	prompt := promptui.Select{
		Label: "Select the Go framework",
		Items: goFrameworks,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("prompt failed %v", err)
	}

	return result, nil
}

func validateProjectName(input string) error {
	if len(input) < 3 {
		return fmt.Errorf("project name must be at least 3 characters")
	}
	return nil
}

func PromptName(label string) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validateProjectName,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed %v", err)
	}
	return result, nil
}

func PromptUsername(label string) (string, error) {
	validate := func(input string) error {
		if len(input) < 10 {
			return errors.New("Username must be at least 10 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed %v", err)
	}
	return result, nil
}

func PromptPassword(label string) (string, error) {
	validate := func(input string) error {
		if len(input) < 16 {
			return errors.New("Password must be at least 16 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Mask:     '*',
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed %v", err)
	}
	return result, nil
}
