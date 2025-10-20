package javascript

import (
	"fmt"

	vueScaffold "github.com/digizyne/lf/internal/languages/javascript/frameworks/vue"
	prompts "github.com/digizyne/lf/internal/prompts"
)

func SelectJavascriptFramework(projectName string) error {
	var selectedFramework string

	selectedFramework, err := prompts.SelectJsFramework()
	if err != nil {
		return fmt.Errorf("prompt failed %v", err)
	}

	switch selectedFramework {
	case "Vue":
		err = vueScaffold.ScaffoldVueProject(projectName)
		if err != nil {
			return fmt.Errorf("vue project scaffolding failed: %v", err)
		}
	}

	return nil
}
