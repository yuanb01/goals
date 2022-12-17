package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yuanb01/goals/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your goals.",
	Run: func(cmd *cobra.Command, args []string) {
		goals, err := db.GetAllGoals()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(goals) == 0 {
			fmt.Println("Your goals list is empty! Why not add a goal? 📝 🥅")
			return
		}
		fmt.Println("You have the following goals:")
		for i, goal := range goals {
			fmt.Printf("%d. %s\n", i+1, goal.Text)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
