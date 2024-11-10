/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
    "fmt"
    // "os"
    "time"

    "github.com/spf13/cobra"
    // "github.com/hardikkum444/go-do-it/models"

)

var todos1 Todos

var addCmd = &cobra.Command{
    Use:   "add [title]",
    Short: "Add a new todo",
    Long:  `Add a new todo item with the specified title.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        title := args[0]
        todos.add(title) 
        fmt.Printf("Added todo: %s\n", title)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}

func(todos *Todos) add(title string) {

    t := *todos

    todo := Todo{
        Title : title,
        Completed : false,
        CreatedAt : time.Now().UTC(),
        CompletedAt : nil,
    }

    t = append(t, todo)
    fmt.Println(todo)
}


