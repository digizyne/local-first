package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	prompts "github.com/digizyne/lf/internal/prompts"
	"github.com/urfave/cli/v3"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(ctx context.Context, cmd *cli.Command) error {
	username, err := prompts.PromptUsername("Enter your lf-cloud username:")
	if err != nil {
		return err
	}

	password, err := prompts.PromptPassword("Enter your lf-cloud password:")
	if err != nil {
		return err
	}

	loginData := LoginRequest{
		Username: username,
		Password: password,
	}
	jsonData, err := json.Marshal(loginData)
	if err != nil {
		return err
	}
	resp, err := http.Post("http://localhost:8080/api/v1/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var loginResp LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configDir := filepath.Join(homeDir, ".config", "lf")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return err
	}
	credentialsFile := filepath.Join(configDir, "credentials.json")
	tokenData := map[string]string{"token": loginResp.Token}
	tokenJSON, err := json.Marshal(tokenData)
	if err != nil {
		return err
	}
	err = os.WriteFile(credentialsFile, tokenJSON, 0600)
	if err != nil {
		return err
	}

	fmt.Println("Login successful. Credentials saved to", credentialsFile)

	return nil
}

func GetBearerToken() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	credentialsFile := filepath.Join(homeDir, ".config", "lf", "credentials.json")
	if _, err := os.Stat(credentialsFile); os.IsNotExist(err) {
		return "", fmt.Errorf("credentials not found, run 'lf login' to reauthenticate")
	}
	credentialsData, err := os.ReadFile(credentialsFile)
	if err != nil {
		return "", err
	}
	var loginResp LoginResponse
	err = json.Unmarshal(credentialsData, &loginResp)
	if err != nil {
		return "", err
	}
	return loginResp.Token, nil
}
