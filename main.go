package main

import (
	"context"
	"log"
	"os"

	"github.com/0p5dev/ops/internal/auth"
	"github.com/0p5dev/ops/internal/deploy"
	"github.com/0p5dev/ops/internal/scaffold"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "ops",
		Usage: "A CLI tool to scaffold and deploy developer-first applications",
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
			{
				Name:    "login",
				Aliases: []string{"l"},
				Usage:   "Login to 0p5.dev",
				Action:  auth.Login,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
