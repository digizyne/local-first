package main

import (
	"context"
	"log"
	"os"

	deploy "github.com/digizyne/lf/internal/deploy"
	scaffold "github.com/digizyne/lf/internal/scaffold"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "lf",
		Usage: "A CLI tool to scaffold and deploy local-first projects",
		Commands: []*cli.Command{
			{
				Name:    "scaffold",
				Aliases: []string{"s"},
				Usage:   "Scaffold a new project",
				Action:  scaffold.Scaffold,
			},
			{
				Name:    "deploy",
				Aliases: []string{"d"},
				Usage:   "Deploy a project",
				Action:  deploy.Deploy,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
