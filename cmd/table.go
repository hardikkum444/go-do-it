/*
Copyright © 2024 man44 <man44@tutamail.com>
*/
package cmd

import (
	"fmt"
	"strconv"
	// "strings"
	"time"

	"github.com/hardikkum444/go-do-it/storage"
	"github.com/rivo/tview"
	// "github.com/gdamore/tcell/v2"
	// "github.com/spf13/cobra"
)

func renderTable() {
	table := tview.NewTable().
		SetBorders(true)

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	err := storage.Load(&todosall)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	table.SetCell(0, 0, tview.NewTableCell("#"))
	table.SetCell(0, 1, tview.NewTableCell("Title"))
	table.SetCell(0, 2, tview.NewTableCell("Completed"))
	table.SetCell(0, 3, tview.NewTableCell("Deadline"))
	table.SetCell(0, 4, tview.NewTableCell("Notes"))
	table.SetCell(0, 5, tview.NewTableCell("CreatedAt"))
	table.SetCell(0, 6, tview.NewTableCell("CompletedAt"))

	for r, todo := range todosall {
		row := r + 1
		table.SetCell(row, 0, tview.NewTableCell(strconv.Itoa(row)))
		table.SetCell(row, 1, tview.NewTableCell(todo.Title))
		table.SetCell(row, 2, tview.NewTableCell(strconv.FormatBool(todo.Completed)))
		table.SetCell(row, 3, tview.NewTableCell(todo.Deadline))
		table.SetCell(row, 4, tview.NewTableCell(todo.Notes))
		table.SetCell(row, 5, tview.NewTableCell((todosall[r].CreatedAt).Format(time.RFC1123)))
		if todo.CompletedAt != nil {
			table.SetCell(row, 6, tview.NewTableCell(todo.CompletedAt.Format(time.RFC1123)))
		} else {
			table.SetCell(row, 6, tview.NewTableCell("null"))
		}
	}

	form := tview.NewForm().
		AddButton("back", func() {
			if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
				panic(err)
			}
		}).
		AddButton("quit", func() {
			renderQuit()
		})

	flex1 := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 3, false).
		AddItem(form, 0, 1, false)

	center := Center(120, 40, flex1)

	if err := app.SetRoot(center, true).EnableMouse(true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
