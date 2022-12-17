package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuanb01/goals/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new goal to your goals list with optional [repeat] param",
	Run: func(cmd *cobra.Command, args []string) {
		var repeat int
		var err error
		if len(args) > 1 { // only want to check if they added optional repeat param if they entered more than 1 argument
			repeat, err = strconv.Atoi(args[len(args)-1])
			if err != nil {
				repeat = 1 // repeat should default to 1 if the goal is not to be repeated
			} else {
				args = args[:len(args)-1] // don't include repeat param in goal text
			}
		}
		goalText := strings.Join(args, " ")
		_, strConvErr := strconv.Atoi(goalText) // goalText cannot solely be a number, otherwise will cause problems
		if strConvErr == nil {
			fmt.Println("Oops, your goal cannot just be a number! Try being a bit more descriptive... 💭")
			return
		}

		goals, getGoalsErr := db.GetAllGoals()
		if getGoalsErr != nil {
			fmt.Println("Something went wrong:", getGoalsErr)
			return
		}
		if goalText != "" {
			for _, goal := range goals {
				if goal.Text == goalText {
					fmt.Printf("Whoops! Looks like \"%s\" is already in your goals list!", goalText)
					return
				}
			}
		} else {
			fmt.Println("✋ Don't forget to write out what goal you would like to add! \nFor more on how to use the 'add' command, try typing 'goals' in your terminal 😉")
			return
		}

		_, createErr := db.CreateGoal(goalText, repeat)
		if createErr != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		if repeat > 1 {
			fmt.Printf("Added \"%s\" to your goals list with repeat %dx!\n", goalText, repeat)
		} else {
			fmt.Printf("Added \"%s\" to your goals list!\n", goalText)
		}
	},
}

func init() { // all init() functions are run before main.go for set up
	RootCmd.AddCommand(addCmd)
}
