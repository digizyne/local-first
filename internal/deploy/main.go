package deploy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/digizyne/lf/internal/auth"
	prompts "github.com/digizyne/lf/internal/prompts"
	"github.com/urfave/cli/v3"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func buildContainerImage(tag string) error {
	cmd := exec.Command("docker", "build", "-t", tag, ".")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func saveContainerImage(tag string, filename string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("docker save %s | gzip > %s", tag, filename))
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func transmitCompressedImage(filename string, token string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/container-registry", file)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/gzip")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if (resp.StatusCode == http.StatusUnauthorized) || (resp.StatusCode == http.StatusForbidden) {
		return fmt.Errorf("authentication failed: please log in again (lf login)")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %v", resp.Status)
	}

	return nil
}

func Deploy(ctx context.Context, cmd *cli.Command) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	credentialsFile := filepath.Join(homeDir, ".config", "lf", "credentials.json")
	if _, err := os.Stat(credentialsFile); os.IsNotExist(err) {
		fmt.Println("No credentials found. Please log in first.")
		err = auth.Login(ctx, cmd)
		if err != nil {
			return fmt.Errorf("login failed: %w", err)
		}
		if _, err := os.Stat(credentialsFile); os.IsNotExist(err) {
			return fmt.Errorf("credentials file still not found after login")
		}
	}
	credentialsData, err := os.ReadFile(credentialsFile)
	if err != nil {
		return fmt.Errorf("failed to read credentials file: %w", err)
	}

	var loginResp LoginResponse
	err = json.Unmarshal(credentialsData, &loginResp)
	if err != nil {
		return fmt.Errorf("failed to parse credentials: %w", err)
	}

	fmt.Printf("Using authentication token: %s...\n", loginResp.Token[:min(8, len(loginResp.Token))])

	deploymentName, err := prompts.PromptName("Deployment Name")
	if err != nil {
		return fmt.Errorf("failed to get deployment name: %w", err)
	}
	filename := fmt.Sprintf("%s.tgz", deploymentName)

	err = buildContainerImage(deploymentName)
	if err != nil {
		return fmt.Errorf("failed to build container image: %v", err)
	}

	err = saveContainerImage(deploymentName, filename)
	if err != nil {
		return fmt.Errorf("failed to save container image: %v", err)
	}

	err = transmitCompressedImage(filename, loginResp.Token)
	if err != nil {
		return fmt.Errorf("failed to transmit compressed image: %v", err)
	}

	fmt.Println("Deployment successful!")

	return nil
}
