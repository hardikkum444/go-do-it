/*
Copyright © 2024 man44 <man44@tutamail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-do-it",
	Short: "to-do application built in go",
    Long: "a simple to-do list application for your cli to make your life more productive and reasonable",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


