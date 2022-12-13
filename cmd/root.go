package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "goals",
	Short: "Goals is a goals tracker right in your terminal",
	// remove Run s.t. default behavior instead of running a command is to show the help text
}

// import (
// 	"fmt"
// 	"os"

// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// )

// var RootCmd = &cobra.Command{
// 	Use:   "goals",
// 	Short: "Goals is a goals tracker right in your terminal",
// 	Long: `A Fast and Flexible Static Site Generator built with
// 				  love by spf13 and friends in Go.
// 				  Complete documentation is available at https://gohugo.io/documentation/`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 	  // Do Stuff Here
// 	},
//   }
  
//   func Execute() {
// 	if err := rootCmd.Execute(); err != nil {
// 	  fmt.Fprintln(os.Stderr, err)
// 	  os.Exit(1)
// 	}
//   }