package main

import (
	"context"
	"fmt"
	"log"
	"os"

	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli-altsrc/v3/yaml"
	"github.com/urfave/cli/v3"
)

func main() {

	// Define a configuration file.
	// Expects "devosync.yml" in the working directory.
	configFile := altsrc.StringSourcer("devosync.yml")

	cmd := &cli.Command{
		Name:  "devosync",
		Usage: "A simple environment sync utility for web developers",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "warn",
				Aliases: []string{"w"},
				Value:   false,
				Usage:   "Warn before performing action.",
				Sources: cli.NewValueSourceChain(yaml.YAML("warn", configFile)),
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "backup",
				Aliases: []string{"b"},
				Usage:   "Backup a database or assets on a remote.",
				Commands: []*cli.Command{
					{
						Name:    "database",
						Aliases: []string{"db", "d"},
						Usage:   "Backup a database on a remote",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "fetch",
								Aliases: []string{"f"},
								Usage:   "Fetch local copies of the item(s) being backed up.",
								Value:   false,
							},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							if cmd.Bool("fetch") {
								fmt.Println("Backing up database on ", cmd.Args().First(), " and fetching a local copy.")
							} else {
								fmt.Println("Backing up database on ", cmd.Args().First())
							}
							return nil
						},
					},
					{
						Name:    "assets",
						Aliases: []string{"ass", "a"},
						Usage:   "Backup assets on a remote",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Backing up assets on ", cmd.Args().First())
							return nil
						},
					},
				},
			},
			{
				Name:    "import",
				Aliases: []string{"i"},
				Usage:   "Import a database or assets locally.",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("Importing ____ to the local environment", cmd.Args().First())
					return nil
				},
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
