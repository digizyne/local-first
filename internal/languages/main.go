package languages

import (
	"fmt"

	jsLanguage "github.com/digizyne/local-first/internal/languages/javascript"
	prompts "github.com/digizyne/local-first/internal/prompts"
)

func SelectProgrammingLanguage(projectName string) error {
	selectedProgrammingLanguage, err := prompts.SelectProgrammingLanguage()
	if err != nil {
		return fmt.Errorf("Prompt failed %v", err)
	}

	switch selectedProgrammingLanguage {
	case "Javascript":
		err = jsLanguage.SelectJavascriptFramework(projectName)
		if err != nil {
			return fmt.Errorf("Javascript framework selection failed %v", err)
		}

	case "Go":
		prompts.SelectGoFramework()
	}

	return nil
}
