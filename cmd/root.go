package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "goals",
	Short: "Goals is a goals tracker right in your terminal",
	// remove Run s.t. default behavior instead of running a command is to show the help text
}