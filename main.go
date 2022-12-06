// package main

// import (
//   "goals/cmd"
// )

// func main() {
//   cmd.RootCmd.Execute()
// }
package main

import (
  "fmt"
  "strings"

  "github.com/spf13/cobra"
)

func main() {
  var echoTimes int
  var rootCmd = &cobra.Command{
    Use:   "goals",
    Short: "Goals is a goals tracker right in your terminal",
  }

  var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new goal to your goals with optional [repeat] param",
    // Args: cobra.MinimumNArgs(1),
    // Run: func(cmd *cobra.Command, args []string) {
    //   fmt.Println("Print: " + strings.Join(args, " "))
    // },
  }

  var doCmd = &cobra.Command{
    Use:   "do",
    Short: "Mark a goal on your goals list as complete by its name or number in the goals list",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Echo: " + strings.Join(args, " "))
    },
  }

  var removeCmd = &cobra.Command{
    Use:   "remove",
    Short: "Remove a goal from your goals list by its name or number in the goals list",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      for i := 0; i < echoTimes; i++ {
        fmt.Println("Echo: " + strings.Join(args, " "))
      }
    },
  }

  var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all of your goals",
    Run: func(cmd *cobra.Command, args []string) {
      for i := 0; i < echoTimes; i++ {
        fmt.Println("Echo: " + strings.Join(args, " "))
      }
    },
  }

  // cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

  rootCmd.AddCommand(addCmd, doCmd, removeCmd, listCmd)
  // cmdEcho.AddCommand(cmdTimes)
  rootCmd.Execute()
}