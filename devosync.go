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
		Name:  "devosync",
		Usage: "A simple environment sync utility for web developers",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "warn",
				Aliases: []string{"w"},
				Value:   false,
				Usage:   "Warn before performing action.",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {

			temp := "Test"

			// Check for Arguments.
			if cmd.NArg() > 0 {
				temp = cmd.Args().Get(0)
			}

			// Warning example that doesn't actually let you cancel.
			if cmd.Bool("warn") {
				fmt.Println("Warning, I'm going to echo some text at you. Unfortunately, there is nothing you can do to stop it.")
				fmt.Println(temp)
			} else {
				fmt.Println(temp)
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
