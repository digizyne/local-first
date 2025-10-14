package selects

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func SelectProgrammingLanguage(items []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select the programming language",
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("Prompt failed %v", err)
	}

	return result, nil
}

func SelectJsFramework(items []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select the Javascript framework",
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("Prompt failed %v", err)
	}

	return result, nil
}

func SelectGoFramework(items []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select the Go framework",
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("Prompt failed %v", err)
	}

	return result, nil
}

func validateProjectName(input string) error {
	if len(input) < 3 {
		return fmt.Errorf("Project name must be at least 3 characters")
	}
	return nil
}

func PromptProjectName() (string, error) {
	prompt := promptui.Prompt{
		Label:    "Project Name",
		Validate: validateProjectName,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("Prompt failed %v", err)
	}
	return result, nil
}
