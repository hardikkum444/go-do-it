/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hardikkum444/go-do-it/storage"
	"github.com/spf13/cobra"
)

var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "toggle completed/not-completed",
	Long:  "toggle whether a task is completed in the todo or not",
	Run: func(cmd *cobra.Command, args []string) {
		index := args[0]
		indexInt, _ := strconv.Atoi(index)

		if err := toggle(indexInt); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("toggled task")
		}
	},
}

func toggle(index int) error {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	if err := validateIndex(index); err != nil {
		return err
	}

	isCompleted := (todosall)[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		todosall[index].CompletedAt = &completionTime
	}

	todosall[index].Completed = !isCompleted

	storage.Save(todosall)
	return nil
}
