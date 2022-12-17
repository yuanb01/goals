package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuanb01/goals/db"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a goal on your goals list as complete either by its name or its number in the goals list",
	Run: func(cmd *cobra.Command, args []string) {
		var goalTextEntered bool
		var goalText string
		arg := strings.Join(args, " ")
		id, err := strconv.Atoi(arg) // id will be 0 if a string is entered
		if err != nil {
			goalText = arg
			goalTextEntered = true
		}
		goals, getGoalsErr := db.GetAllGoals()
		if getGoalsErr != nil {
			fmt.Println("Something went wrong:", getGoalsErr)
			return
		}
		if goalText == "" { // means user entered empty string for do
			fmt.Println("✋ Don't forget to write out what goal you would like to mark as done! \nFor more on how to use the 'do' command, try typing 'goals' in your terminal 😉")
			return
		} else {
			for i, goal := range goals {
				if goal.Text == goalText {
					id = i + 1
				}
			}
		}

		if goalTextEntered && id <= 0 {
			fmt.Printf("Uh oh! Did you make a typo? \"%s\" doesn't seem to be in your goals list.. 🧐\n📝 Use the 'goals list' command to see your list of goals!\n", goalText)
			return
		}
		if id > len(goals) {
			fmt.Printf("Uh oh! Did you make a typo? There is no goal at #%d in your goals list.. 🧐\n📝 Use the 'goals list' command to see your list of goals!\n", id)
			return
		}
		goal := goals[id-1]
		if goal.Repeat <= 1 {
			deleteErr := db.DeleteGoal(goal.Id)
			if deleteErr != nil {
				fmt.Printf("Failed to mark your goal \"%s\" as completed. Error: %s\n", goal.Text, err)
			} else {
				fmt.Printf("Yay! You have completed your \"%s\" goal! 🎉\n", goal.Text)
			}
		} else {
			updateErr := db.UpdateGoal(goal.Id, goal.Text, goal.Repeat-1)
			if updateErr != nil {
				fmt.Printf("Failed to mark your goal \"%s\" as completed. Error: %s\n", goal.Text, err)
			} else {
				fmt.Printf("You are making progress towards your \"%s\" goal! You need to do this goal %d more times 🔁\n", goal.Text, goal.Repeat-1)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
