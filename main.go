package main

import (
	"fmt"

	languages "github.com/digizyne/local-first/internal/languages"
	prompts "github.com/digizyne/local-first/internal/prompts"
)

func main() {
	projectName, err := prompts.PromptProjectName()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	err = languages.SelectProgrammingLanguage(projectName)
	if err != nil {
		fmt.Printf("Programming language selection failed %v\n", err)
		return
	}

	fmt.Println("Project scaffolding completed successfully.")
}
