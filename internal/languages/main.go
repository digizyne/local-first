package languages

import (
	"fmt"

	jsLanguage "github.com/0p5dev/ops/internal/languages/javascript"
	prompts "github.com/0p5dev/ops/internal/prompts"
)

func SelectProgrammingLanguage(projectName string) error {
	selectedProgrammingLanguage, err := prompts.SelectProgrammingLanguage()
	if err != nil {
		return fmt.Errorf("prompt failed %v", err)
	}

	switch selectedProgrammingLanguage {
	case "Javascript":
		err = jsLanguage.SelectJavascriptFramework(projectName)
		if err != nil {
			return fmt.Errorf("javascript framework selection failed %v", err)
		}

	case "Go":
		prompts.SelectGoFramework()
	}

	return nil
}
