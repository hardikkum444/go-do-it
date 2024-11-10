/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/hardikkum444/go-do-it/storage"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit an existing task",
	Long:  "edit/correct an existing task in your todo",
	Run: func(cmd *cobra.Command, args []string) {

		index := args[0]
		title := args[1]
		indexInt, _ := strconv.Atoi(index)

		if err := edit(indexInt, title); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("task edited")
		}
	},
}

func edit(index int, title string) error {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	if err := validateIndex(index); err != nil {
		return err
	}

	todosall[index].Title = title

	storage.Save(todosall)
	return nil
}
