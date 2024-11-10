/*
Copyright © 2024 man44 <man44@tutamail.com>

*/

package cmd

import (
    "fmt"
    // "os"
    "time"

    "github.com/spf13/cobra"
    // "github.com/hardikkum444/go-do-it/models"

)

var addCmd = &cobra.Command{
    Use:   "add [title]",
    Short: "Add a new todo",
    Long:  `Add a new todo item with the specified title.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        title := args[0]
        todos.add(title) 
        fmt.Printf("Added todo: %s\n", title)
        fmt.Println("Current todos:", todos)
    },
}

// func init() {
//     rootCmd.AddCommand(addCmd)
// }

func(todos *Todos) add(title string) {

    todo := Todo{
        Title : title,
        Completed : false,
        CreatedAt : time.Now().UTC(),
        CompletedAt : nil,
    }

    *todos = append(*todos, todo)
    // fmt.Println(*todos)
}


