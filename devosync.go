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
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("hey girl")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
