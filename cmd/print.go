/*
Copyright © 2024 man44 <man44@tutamail.com>
*/
package cmd

import (
	"github.com/aquasecurity/table"
	"github.com/hardikkum444/go-do-it/storage"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print the todo table",
	Long:  `printing all the contents of the todo table which has been prepared`,
	Run: func(cmd *cobra.Command, args []string) {
		todos.print()
	},
}

func (todos *Todos) print() {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Deadline", "Notes", "Completed", "Created At", "Completed At")

	for index, t := range todosall {
		completed := "❌"
		completedAt := "-"

		if t.Deadline == "" {
			t.Deadline = "-"
		}

		if t.Notes == "" {
			t.Notes = "-"
		}

		if t.Completed {
			completed = "✔️"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, t.Deadline, t.Notes, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
