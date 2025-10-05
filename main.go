package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "init",
		Usage: "Create initial archive.yml.",
		Action: func(context.Context, *cli.Command) error {
			err := initConfig()
			if err != nil {
				return fmt.Errorf("Failed to initialize config: %w", err)
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
