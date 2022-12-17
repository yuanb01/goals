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
	Short: "Adds a goal to your goals list.",
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
		_, createErr := db.CreateGoal(goalText, repeat)
		if createErr != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("Added \"%s\" to your goals list.\n", goalText)
	},
}

func init() { // all init() functions are run before main.go for set up
	RootCmd.AddCommand(addCmd)
}
