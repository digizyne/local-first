package deploy

import (
	"context"
	"fmt"
	"net/http"
	"os"
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

func saveContainerImage(tag string, filename string) error {
	cmd := exec.Command("docker", "save", tag, "|", "gzip", ">", filename)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("docker save failed: %v", err)
	}

	return nil
}

func transmitCompressedImage(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filename, err)
	}
	defer file.Close()

	req, err := http.NewRequest("POST", "https://example.com/upload", file)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/gzip")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to transmit image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to transmit image: %v", resp.Status)
	}

	return nil
}

func Deploy(ctx context.Context, cmd *cli.Command) error {
	deploymentName, err := prompts.PromptName("Deployment Name")
	if err != nil {
		return fmt.Errorf("failed to get deployment name: %w", err)
	}
	filename := fmt.Sprintf("%s.tgz", deploymentName)

	err = BuildContainerImage(deploymentName)
	if err != nil {
		return fmt.Errorf("failed to build container image: %v", err)
	}

	err = saveContainerImage(deploymentName, filename)
	if err != nil {
		return fmt.Errorf("failed to save container image: %v", err)
	}

	err = transmitCompressedImage(filename)
	if err != nil {
		return fmt.Errorf("failed to transmit compressed image: %v", err)
	}

	return nil
}
