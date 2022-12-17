package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/yuanb01/goals/cmd"
	"github.com/yuanb01/goals/db"
)

func main() {
	home, _ := homedir.Dir() // get home directory for a user
	dbPath := filepath.Join(home, "goals.db")
	err := db.Init(dbPath)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		os.Exit(1)
	}
	cmd.RootCmd.Execute()
}
