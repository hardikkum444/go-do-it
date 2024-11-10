/*
Copyright © 2024 man44 <man44@tutamail.com>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/hardikkum444/go-do-it/storage"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "deleting a task",
	Long:  "this command will delete a task present in your todolist",
	Run: func(cmd *cobra.Command, args []string) {
		index := args[0]
		indexInt, _ := strconv.Atoi(index)
		err := del(indexInt)
		if err != nil {
			fmt.Println(err, indexInt)
		} else {
			fmt.Println("indexed task has been deleted")
		}
	},
}

func del(index int) error {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	if err := validateIndex(index); err != nil {
		return err
	}

	todosall = append(todosall[:index], todosall[index+1:]...)
	storage.Save(todosall)

	return nil
}
