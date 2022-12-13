package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your goals.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}