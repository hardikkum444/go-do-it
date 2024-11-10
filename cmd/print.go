/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "fmt"
    "github.com/aquasecurity/table" 
    "time"
	"github.com/spf13/cobra"
    "os"
    "strconv"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print the todo table",
	Long:   `printing all the contents of the todo table which has been prepared`, 
	Run: func(cmd *cobra.Command, args []string) {
        todos.print()
	},
}

func(todos *Todos) print() {

    table := table.New(os.Stdout)
    table.SetRowLines(false)
    table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

    for index, t := range *todos {
        completed := "❌"
        completedAt := ""

        if t.Completed {
            completed = "✅"
            if t.CompletedAt != nil{
                completedAt = t.CompletedAt.Format(time.RFC1123)

            }
        }

        table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
    }

    table.Render()

}
