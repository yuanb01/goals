package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a goal to your goals list.",
	Run: func(cmd *cobra.Command, args []string) {
		goal := strings.Join(args, " ") // lets user enter do dishes or "do dishes"
		fmt.Printf("Added \"%s\" to your goals list.\n", goal)
	},
}

func init() { // all init() functions are run before main.go for set up 
	RootCmd.AddCommand(addCmd)
}