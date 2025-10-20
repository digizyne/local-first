package deploy

import (
	"context"
	"fmt"
	"os/exec"

	prompts "github.com/digizyne/lf/internal/prompts"
	"github.com/urfave/cli/v3"
)

func BuildContainerImage(tag string) error {
	cmd := exec.Command("docker", "build", "-t", tag, ".")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("docker build failed: %v", err)
	}

	return nil
}

func PushContainerImage(tag string) error {
	cmd := exec.Command("docker", "push", tag)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("docker push failed: %v", err)
	}

	return nil
}

func Deploy(ctx context.Context, cmd *cli.Command) error {
	deploymentName, err := prompts.PromptName("Deployment Name")
	if err != nil {
		return fmt.Errorf("failed to get deployment name: %w", err)
	}

	tag := "us-central1-docker.pkg.dev/jcastle-dev/local-first-public/" + deploymentName + ":latest"

	err = BuildContainerImage(tag)
	if err != nil {
		return fmt.Errorf("failed to build container image: %v", err)
	}

	err = PushContainerImage(tag)
	if err != nil {
		return fmt.Errorf("failed to push container image: %v", err)
	}

	return nil
}
