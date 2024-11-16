/*
Copyright © 2024 man44 <man44@tutamail.com>
*/
package cmd

import (
	// "fmt"
	"time"

	"github.com/hardikkum444/go-do-it/storage"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "open go-do-it tui",
	Long:  "open the terminal user interface",
	Run: func(cmd *cobra.Command, args []string) {
		renderMenu()
	},
}

var (
	app  *tview.Application
	list *tview.List
	form *tview.Form
)

func createMenuList(app *tview.Application) *tview.List {

	list = tview.NewList().
		AddItem("add", "add a new task", 'n', func() {
			renderAdd(app)
		}).
		AddItem("quit", "quit application", 'q', func() {
			renderQuit()
		})

	return list
}

func renderMenu() {

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
			addToTable(taskTitle.GetText(), taskDeadline.GetText(), taskNotes.GetText())
			renderDone()
		}).
		AddButton("quit", func() {
			renderQuit()
		})

	form.SetBorder(true).SetTitle("add a task").SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func addToTable(title string, deadline string, notes string) {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	todo := Todo{
		Title:       title,
		Deadline:    deadline,
		Notes:       notes,
		Completed:   false,
		CreatedAt:   time.Now().UTC(),
		CompletedAt: nil,
	}

	todosall = append(todosall, todo)
	storage.Save(todosall)

}

func renderQuit() {

	modal := tview.NewModal().
		SetText("do you want to exit tui?").
		AddButtons([]string{"cancle", "quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "quit" {
				app.Stop()
			} else if buttonLabel == "cancle" {
				if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
					panic(err)
				}
			}
		})

	if err := app.SetRoot(modal, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func renderDone() {

	modal := tview.NewModal().
		SetText("successful").
		AddButtons([]string{"ok"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "ok" {
				if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
					panic(err)
				}
			}
		})

	if err := app.SetRoot(modal, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
