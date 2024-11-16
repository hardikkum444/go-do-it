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
	Short: "add a new todo",
	Long:  `Add a new todo item with the specified title.`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		deadline := args[1]
        notes := args[2]
		add(title, deadline, notes)
		fmt.Printf("Added todo: %s\n", title)
	},
}

func add(title string, deadline string, notes string) {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	todo := Todo{
		Title:       title,
        Deadline:    deadline,
        Notes:       notes,
		Completed:   false,
		CreatedAt:   time.Now().UTC(),
		CompletedAt: nil,
	}

	todosall = append(todosall, todo)
	storage.Save(todosall)
}
