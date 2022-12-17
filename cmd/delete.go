package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuanb01/goals/db"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a goal from your goals list by its name or its number in the goals list",
	Run: func(cmd *cobra.Command, args []string) {
		var goalText string
		arg := strings.Join(args, " ")
		id, err := strconv.Atoi(arg) // id will be 0 if a string is entered
		if err != nil {              // referring to goal by string name
			goalText = arg
		}
		goals, getGoalsErr := db.GetAllGoals()
		if getGoalsErr != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		if goalText != "" {
			for i, goal := range goals {
				if goal.Text == goalText {
					id = i + 1
				}
			}
		}

		if id <= 0 {
			fmt.Printf("Uh oh! Did you make a typo? \"%s\" doesn't seem to be in your goals list.. 🧐\n", goalText)
			return
		}
		if id > len(goals) {
			fmt.Printf("Uh oh! Did you make a typo? There is no goal at #%d in your goals list.. 🧐\n", id)
			return
		}

		goal := goals[id-1]
		deleteErr := db.DeleteGoal(goal.Id)
		if deleteErr != nil {
			fmt.Printf("Failed to delete your goal \"%s\". Error: %s\n", goal.Text, err)
		} else {
			fmt.Printf("Deleted goal \"%s\" 🗑\n", goal.Text)
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
