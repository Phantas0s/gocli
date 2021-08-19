package cmd

import (
	"github.com/Phantas0s/go-cli/internal"
	"github.com/spf13/cobra"
)

var archive bool

func otherCmd() *cobra.Command {
	otherCmd := &cobra.Command{
		Use:   "list",
		Short: "List your Pocket pages",
		// TODO write some help
		Run: func(cmd *cobra.Command, args []string) {
			runList()
		},
	}

	otherCmd.Flags().BoolVarP(&archive, "archive", "a", false, "Archive the listed articles (with confirmation)")

	return otherCmd
}

func runOther() {
	internal.Display()
}
