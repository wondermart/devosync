package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	/* 	Re: "Organize Commands in Packages"
	https://cobra.dev/docs/how-to-guides/working-with-commands/

	This would likely be an ideal approach, to allow separation between the
	CLI structure and the program itself.

	BUT, does that mean I need separate repos for each package, like the
	example shows?

	*/)

var (
// init flag vars?
// why is this () not {}?
)

var rootCmd = &cobra.Command{
	Use:   "devosync [tbd] [tbd]",
	Short: "A simple environment sync utility for web developers",
	Long: `Devosync is a CLI utility for easily syncing environments for web development.

	Add more here later about how devosync can backup, fetch, and import databases or files. `,
	// tk_args - can't find a docs page that explains this field.
	Run: runDevosync,
}

func init() {
	// tk_define_flags
}

func runDevosync(cmd *cobra.Command, args []string) {
	// fmt.Println(cmd)
	fmt.Println("hey girl")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
