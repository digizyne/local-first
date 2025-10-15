package javascript

import (
	"fmt"

	vueScaffold "github.com/digizyne/local-first/internal/languages/javascript/frameworks/vue"
	prompts "github.com/digizyne/local-first/internal/prompts"
)

func SelectJavascriptFramework(projectName string) error {
	var selectedFramework string

	selectedFramework, err := prompts.SelectJsFramework()
	if err != nil {
		return fmt.Errorf("Prompt failed %v", err)
	}

	switch selectedFramework {
	case "Vue":
		err = vueScaffold.ScaffoldVueProject(projectName)
		if err != nil {
			return fmt.Errorf("Vue project scaffolding failed: %v", err)
		}
	}

	return nil
}
