/*
Copyright © 2024 man44 <man44@tutamail.com>
*/
package cmd

import (
	"errors"
	"os"
	"time"

	"github.com/hardikkum444/go-do-it/storage"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-do-it",
	Short: "to-do application built in go",
	Long:  "a simple to-do list application for your cli to make your life more productive and reasonable",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

var todos Todos

func validateIndex(index int) error {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	if index < 0 || index >= len(todosall) {
		return errors.New("invalid index")
	}
	return nil
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(toggleCmd)
	rootCmd.AddCommand(editCmd)

}
