package main

import (
	"fmt"
	"os/exec"

	dockerfiles "github.com/digizyne/local-first/internal/dockerfiles"
	prompts "github.com/digizyne/local-first/internal/prompts"
)

func main() {
	programmingLanguages := []string{"Javascript", "Go"}
	jsFrameworks := []string{"React", "Vue", "Angular"}
	goFrameworks := []string{"Gin", "Echo", "Fiber"}

	selectedProgrammingLanguage, err := prompts.SelectProgrammingLanguage(programmingLanguages)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	var selectedJsFramework string

	switch selectedProgrammingLanguage {
	case "Javascript":
		selectedJsFramework, err = prompts.SelectJsFramework(jsFrameworks)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

	case "Go":
		prompts.SelectGoFramework(goFrameworks)
	}

	projectName, err := prompts.PromptProjectName()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch selectedJsFramework {
	case "Vue":
		cmd := exec.Command("npm", "create", "vue@latest", "--", "--default", projectName)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Command execution failed: %v\n", err)
			return
		}

		err = dockerfiles.CreateDockerfile()
		if err != nil {
			fmt.Printf("Error creating Dockerfile: %v\n", err)
			return
		}

		err = dockerfiles.CreateDockerComposeFile()
		if err != nil {
			fmt.Printf("Error creating docker-compose.yml: %v\n", err)
			return
		}
	}
}
