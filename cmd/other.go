package cmd

import (
	"github.com/Phantas0s/gocli/internal"
	"github.com/spf13/cobra"
)

var archive bool

func otherCmd() *cobra.Command {
	otherCmd := &cobra.Command{
		Use:   "example",
		Short: "Do something",
		Run: func(cmd *cobra.Command, args []string) {
			doSomething()
		},
	}

	otherCmd.Flags().BoolVarP(&archive, "archive", "a", false, "Archive the listed articles (with confirmation)")

	return otherCmd
}

func doSomething() {
	internal.Display()
}
