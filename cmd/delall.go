/*
Copyright © 2024 man44 <man44@tutamail.com>
*/
package cmd

import (
	"fmt"

	"github.com/hardikkum444/go-do-it/storage"
	"github.com/spf13/cobra"
)

var delallCmd = &cobra.Command{
	Use:   "delall",
	Short: "deleting all items of the list",
	Long:  "delall will delete all items present in your todo list",
	Run: func(cmd *cobra.Command, args []string) {
        delall()
        fmt.Println("deleted all items")
	},
}

func delall() {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	todosall = Todos{}

	storage.Save(todosall)

}
