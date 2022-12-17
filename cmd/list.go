package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yuanb01/goals/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your goals.",
	Run: func(cmd *cobra.Command, args []string) {
		// db.DeleteGoal(4) // debugging a problem
		goals, err := db.GetAllGoals()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		if len(goals) == 0 {
			fmt.Println("Your goals list is empty! Why not add a goal? 📝 🥅")
			return
		}
		fmt.Println("You have the following goals:")
		for i, goal := range goals {
			if goal.Repeat > 1 {
				fmt.Printf("%d. %s %d\n", i+1, goal.Text, goal.Repeat)
			} else {
				fmt.Printf("%d. %s\n", i+1, goal.Text)
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
