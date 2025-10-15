package deploy

import (
	"context"
	"fmt"
	"os/exec"

	prompts "github.com/digizyne/local-first/internal/prompts"
	"github.com/urfave/cli/v3"
)

func BuildContainerImage(containerImageName string) error {
	tag := "us-central1-docker.pkg.dev/jcastle-dev/local-first-public/" + containerImageName
	cmd := exec.Command("docker", "build", "-t", tag, ".")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("docker build failed: %v", err)
	}

	return nil
}

func PushContainerImage(containerImageName string) error {
	cmd := exec.Command("docker", "push", containerImageName)
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

	err = BuildContainerImage(deploymentName)
	if err != nil {
		return fmt.Errorf("failed to build container image: %v", err)
	}

	err = PushContainerImage(deploymentName)
	if err != nil {
		return fmt.Errorf("failed to push container image: %v", err)
	}

	return nil
}
