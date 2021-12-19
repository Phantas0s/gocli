package cmd

import (
	"{{.VCS}}/{{.User}}/{{.Name}}/internal"
	"github.com/spf13/cobra"
)

type exampleFlags struct {
	example bool
}

func exampleCmd() *cobra.Command {
	exampleFlags := exampleFlags{}

	exampleCmd := &cobra.Command{
		Use:   "example",
		Short: "Do something",
		Run: func(cmd *cobra.Command, args []string) {
			doSomething(exampleFlags)
		},
	}

	exampleCmd.Flags().BoolVarP(&exampleFlags.example, "example", "e", false, "This is an example flag!")

	return exampleCmd
}

func doSomething(exampleFlags exampleFlags) {
	internal.Display(exampleFlags.example)
}
