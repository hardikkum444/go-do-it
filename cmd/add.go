/*
Copyright © 2024 man44 <man44@tutamail.com>

*/

package cmd

import (
	"fmt"
	"github.com/hardikkum444/go-do-it/storage"
	"github.com/spf13/cobra"
	"time"
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new todo",
	Long:  `Add a new todo item with the specified title.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		add(title)
		fmt.Printf("Added todo: %s\n", title)
	},
}

func add(title string) {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now().UTC(),
		CompletedAt: nil,
	}

	todosall = append(todosall, todo)
	storage.Save(todosall)
}
