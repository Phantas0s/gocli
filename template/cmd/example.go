package cmd

import (
	"github.com/Phantas0s/{{.Name}}/internal"
	"github.com/spf13/cobra"
)

type ExampleFlags struct {
	example bool
}

func exampleCmd() *cobra.Command {
	exampleFlags := ExampleFlags{}

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

func doSomething(exampleFlags ExampleFlags) {
	internal.Display(exampleFlags.example)
}
