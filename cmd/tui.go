/*
Copyright © 2024 man44 <man44@tutamail.com>
*/
package cmd

import (
	// "fmt"
	"strconv"
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

func createMenuList() *tview.List {

	list = tview.NewList().
		AddItem("add", "add a new task", 'a', func() {
			renderAdd()
		}).
		AddItem("edit", "edit a new task", 'e', func() {
            renderEdit()
		}).
		AddItem("delete", "delete a task", 'd', func() {
			renderDel()
		}).
		AddItem("delete all", "delete all tasks", 'x', func() {
			renderDelall()
		}).
		AddItem("quit", "quit application", 'q', func() {
			renderQuit()
		})

	return list
}

func renderMenu() {

	app = tview.NewApplication()

	list = createMenuList()

	if err := app.SetRoot(list, true).EnableMouse(true).SetFocus(list).Run(); err != nil {
		panic(err)
	}

}

func renderAdd() {

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
		AddButton("back", func() {
			if err := app.SetRoot(list, true).EnableMouse(true).SetFocus(list).Run(); err != nil {
				panic(err)
			}
		}).
		AddButton("quit", func() {
			renderQuit()
		})

	form.SetBorder(true).SetTitle(" add a task ").SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func renderEdit() {

    taskTitle := tview.NewInputField().SetLabel("task -> ").SetFieldWidth(20)
	taskDeadline := tview.NewInputField().SetLabel("deadline -> ").SetFieldWidth(20)
	taskNotes := tview.NewInputField().SetLabel("notes -> ").SetFieldWidth(20)

    storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	taskIndexes := []string{}
	for index, _ := range todosall {
		taskIndexes = append(taskIndexes, strconv.Itoa(index))
	}

	taskIndex := tview.NewDropDown().
		SetLabel("select task index to edit (hit enter): ").
		SetOptions(taskIndexes, nil)

    form := tview.NewForm().
    AddFormItem(taskIndex).
    AddFormItem(taskTitle).
    AddFormItem(taskDeadline).
    AddFormItem(taskNotes).
    AddButton("done", func() {

    }).
    AddButton("back", func() {
		if err := app.SetRoot(list, true).EnableMouse(true).SetFocus(list).Run(); err != nil {
			panic(err)
		}
	}).
	AddButton("quit", func() {
		renderQuit()
	})

    form.SetBorder(true).SetTitle(" edit task ").SetTitleAlign(tview.AlignCenter)

    if err := app.SetRoot(form, true).EnableMouse(true).SetFocus(form).Run(); err != nil {
        panic(err)
    }
}

func renderDel() {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	taskIndexes := []string{}
	for index, _ := range todosall {
		taskIndexes = append(taskIndexes, strconv.Itoa(index))
	}

	taskIndex := tview.NewDropDown().
		SetLabel("select an index (hit enter): ").
		SetOptions(taskIndexes, nil)

	form = tview.NewForm().
		AddFormItem(taskIndex).
		AddButton("del", func() {
			_, option := taskIndex.GetCurrentOption()
			indexToDel, _ := strconv.Atoi(option)
			delFromTable(indexToDel)
			renderDone()
		}).
		AddButton("back", func() {
			if err := app.SetRoot(list, true).EnableMouse(true).SetFocus(list).Run(); err != nil {
				panic(err)
			}
		}).
		AddButton("quit", func() {
			renderQuit()
		})

	form.SetBorder(true).SetTitle(" delete a task ").SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func renderDelall() {

	modal := tview.NewModal().
		SetText("delete all items in todo list").
		AddButtons([]string{"delete", "cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "delete" {
				delallFromTable()
				renderDone()
			} else if buttonLabel == "cancel" {
				if err := app.SetRoot(list, true).EnableMouse(true).SetFocus(list).Run(); err != nil {
					panic(err)
				}
			}
		})

	if err := app.SetRoot(modal, true).EnableMouse(true).SetFocus(modal).Run(); err != nil {
		panic(err)
	}

}

func addToTable(title string, deadline string, notes string) {

	if title == "" {

		modal := tview.NewModal().
			SetText("title cannot be empty").
			AddButtons([]string{"ok"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "ok" {
					if err := app.SetRoot(list, true).EnableMouse(true).SetFocus(list).Run(); err != nil {
						panic(err)
					}
				}
			})

		if err := app.SetRoot(modal, true).EnableMouse(true).SetFocus(modal).Run(); err != nil {
			panic(err)
		}

	}

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

func delFromTable(index int) {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	todosall = append(todosall[:index], todosall[index+1:]...)
	storage.Save(todosall)

}

func delallFromTable() {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	todosall = Todos{}
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
				if err := app.SetRoot(list, true).EnableMouse(true).SetFocus(list).Run(); err != nil {
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
		SetText("successful!").
		AddButtons([]string{"ok"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "ok" {
				if err := app.SetRoot(list, true).EnableMouse(true).SetFocus(list).Run(); err != nil {
					panic(err)
				}
			}
		})

	if err := app.SetRoot(modal, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
