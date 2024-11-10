/*
Copyright © 2024 man44 <man44@tutamail.com>

*/
package cmd

import (
	"os"
    "time"
    "errors"

	"github.com/spf13/cobra"
    "github.com/hardikkum444/go-do-it/storage"

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

type Todo struct{
    Title string
    Completed bool
    CreatedAt time.Time
    CompletedAt *time.Time
}

type Todos []Todo 

var todos Todos

// var Storage *storage.Storage[Todos]

func validateIndex(index int) error {

    storage := storage.NewStorage[Todos]("todos.json")
    todosall := Todos{}
    storage.Load(&todosall)

    if index < 0 || index >= len(todosall){
        return errors.New("invalid index")
    }
    return nil
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(printCmd)
    // rootCmd.AddCommand(listCmd)
    rootCmd.AddCommand(deleteCmd)
    // rootCmd.AddCommand(toggleCmd)

}
