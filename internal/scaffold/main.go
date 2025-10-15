package scaffold

import (
	"context"
	"fmt"

	languages "github.com/digizyne/local-first/internal/languages"
	prompts "github.com/digizyne/local-first/internal/prompts"
	"github.com/urfave/cli/v3"
)

func Scaffold(ctx context.Context, cmd *cli.Command) error {
	projectName, err := prompts.PromptProjectName()
	if err != nil {
		return fmt.Errorf("failed to get project name: %w", err)
	}

	err = languages.SelectProgrammingLanguage(projectName)
	if err != nil {
		return fmt.Errorf("failed to select programming language: %w", err)
	}

	fmt.Println("Project scaffolding completed successfully.")
	return nil
}
