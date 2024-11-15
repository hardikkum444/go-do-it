/*
Copyright © 2024 man44 <man44@tutamail.com>
*/
package cmd

import (
	// "fmt"

	"github.com/hardikkum444/go-do-it/storage"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"time"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "open go-do-it tui",
	Long:  "open the terminal user interface",
	Run: func(cmd *cobra.Command, args []string) {
		menu()
	},
}

var (
	app  *tview.Application
	list *tview.List
	form *tview.Form
)

func createMenuList(app *tview.Application) *tview.List {

	list = tview.NewList().
		AddItem("add a task", "", 'n', func() {
			renderAdd(app)
		}).
		AddItem("quit", "", 'q', func() {
			app.Stop()
		})

	return list
}

func menu() {

	app = tview.NewApplication()

	list = createMenuList(app)

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}

}

func renderAdd(app *tview.Application) {

	taskTitle := tview.NewInputField().SetLabel("task -> ").SetFieldWidth(20)
	taskDeadline := tview.NewInputField().SetLabel("deadline -> ").SetFieldWidth(20)
	taskNotes := tview.NewInputField().SetLabel("notes -> ").SetFieldWidth(20)

	form := tview.NewForm().
		AddFormItem(taskTitle).
		AddFormItem(taskDeadline).
		AddFormItem(taskNotes).
		AddButton("add", func() {
			sendData(taskTitle.GetText())
		}).
		AddButton("quit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("add a task").SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(form, true).Run(); err != nil {
		panic(err)
	}

}

func sendData(title string) {
	addToTable(title)
}

func addToTable(title string) {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now().UTC(),
		CompletedAt: nil,
	}

	todosall = append(todosall, todo)
	storage.Save(todosall)

}
