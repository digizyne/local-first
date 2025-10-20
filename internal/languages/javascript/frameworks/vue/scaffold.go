package vue

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func createFiles(projectName string) error {
	files := []map[string]string{
		{
			"destination": "./" + projectName + "/Dockerfile",
			"source":      "Dockerfile-vue",
		},
		{
			"destination": "./" + projectName + "/docker-compose.yml",
			"source":      "docker-compose-vue.yml",
		},
		{
			"destination": "./" + projectName + "/.dockerignore",
			"source":      "dockerignore-vue",
		},
		{
			"destination": "./" + projectName + "/nginx.conf",
			"source":      "nginx-vue.conf",
		},
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var success bool
	defer func() {
		if !success && err != nil {
			if cleanupErr := os.RemoveAll("./" + projectName); cleanupErr != nil {
				err = fmt.Errorf("%v; additionally failed to remove directory %s: %v", err, projectName, cleanupErr)
			}
		}
	}()

	for _, f := range files {
		rc, err := client.Bucket("lf-public-scaffold-artifacts").Object(f["source"]).NewReader(ctx)
		if err != nil {
			return fmt.Errorf("Object(%q).NewReader: %w", f["source"], err)
		}
		defer rc.Close()

		localFile, err := os.Create(f["destination"])
		if err != nil {
			return fmt.Errorf("os.Create(%q): %w", localFile.Name(), err)
		}
		defer func() {
			if closeErr := localFile.Close(); closeErr != nil && err == nil {
				err = fmt.Errorf("localFile.Close: %w", closeErr)
			}
		}()

		_, err = io.Copy(localFile, rc)
		if err != nil {
			return fmt.Errorf("io.Copy failed: %w", err)
		}
	}

	success = true

	return nil
}

func ScaffoldVueProject(projectName string) (err error) {
	// Create the Vue project first
	cmd := exec.Command("npm", "create", "vue@latest", "--", "--default", projectName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command execution failed: %v", err)
	}

	// Create additional files
	if err := createFiles(projectName); err != nil {
		return fmt.Errorf("error creating additional files: %v", err)
	}

	return nil
}
