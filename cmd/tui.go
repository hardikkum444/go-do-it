/*
Copyright © 2024 man44 <man44@tutamail.com>
*/
package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/hardikkum444/go-do-it/storage"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "open go-do-it tui",
	Long:  "open the terminal user interface",
	Run: func(cmd *cobra.Command, args []string) {
		renderMenu()
	}}
var (
	app          *tview.Application
	list         *tview.List
	form         *tview.Form
	centeredRoot tview.Primitive
)

func Center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(tview.NewTextView(), 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(tview.NewTextView(), 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(tview.NewTextView(), 0, 1, false), width, 1, true).
		AddItem(tview.NewTextView(), 0, 1, false)
}

func createMenuList() *tview.List {

	list = tview.NewList().
		AddItem("add", "add a new task", 'a', func() {
			renderAdd()
		}).
		AddItem("edit", "edit a new task", 'e', func() {
			renderEdit()
		}).
		AddItem("toggle", "toggle completion of task", 't', func() {
			renderToggle()
		}).
		AddItem("delete", "delete a task", 'd', func() {
			renderDel()
		}).
		AddItem("delete all", "delete all tasks", 'x', func() {
			renderDelall()
		}).
		AddItem("render table", "render tasks table", 'r', func() {
			renderTable()
		}).
		AddItem("quit", "quit application", 'q', func() {
			renderQuit()
		})

	return list
}

func renderMenu() {
	app = tview.NewApplication()

	list = createMenuList()

	centeredRoot = Center(40, 14, list)

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'k' {
			return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
		} else if event.Rune() == 'j' {
			return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
		} else if event.Rune() == 'h' {
			return tcell.NewEventKey(tcell.KeyLeft, 0, tcell.ModNone)
		} else if event.Rune() == 'l' {
			return tcell.NewEventKey(tcell.KeyRight, 0, tcell.ModNone)
		}
		return event
	})

	if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
		panic(err)
	}
}

func renderAdd() {

	taskTitle := tview.NewInputField().SetLabel("Add a task to todo list ").SetFieldWidth(20)
	taskDeadline := tview.NewInputField().SetLabel("Set a completion deadline ").SetFieldWidth(20)
	taskNotes := tview.NewInputField().SetLabel("Add notes ").SetFieldWidth(20)

	form := tview.NewForm().
		AddFormItem(taskTitle).
		AddFormItem(taskDeadline).
		AddFormItem(taskNotes).
		AddButton("add", func() {
			addToTable(taskTitle.GetText(), taskDeadline.GetText(), taskNotes.GetText())
			renderMessage("Task added successfully!")
		}).
		AddButton("back", func() {
			if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
				panic(err)
			}
		}).
		AddButton("quit", func() {
			renderQuit()
		})

	form.SetBorder(false).SetTitle(" add a task ").SetTitleAlign(tview.AlignCenter)

	centeredForm := Center(50, 10, form)

	if err := app.SetRoot(centeredForm, true).EnableMouse(true).SetFocus(centeredForm).Run(); err != nil {
		panic(err)
	}

}

func renderEdit() {

	taskTitle := tview.NewInputField().SetLabel("Task ").SetFieldWidth(20)
	taskDeadline := tview.NewInputField().SetLabel("Deadline ").SetFieldWidth(20)
	taskNotes := tview.NewInputField().SetLabel("Notes ").SetFieldWidth(20)

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	if len(todosall) == 0 {
		renderMessage("Error: Nothing to edit, please add a task!")
	}

	taskIndexes := []string{}
	for index, _ := range todosall {
		taskIndexes = append(taskIndexes, strconv.Itoa(index))
	}

	taskIndex := tview.NewDropDown().
		SetLabel("Select task index to edit (hit enter): ").
		SetOptions(taskIndexes, nil)

	form := tview.NewForm().
		AddFormItem(taskIndex).
		AddFormItem(taskTitle).
		AddFormItem(taskDeadline).
		AddFormItem(taskNotes).
		AddButton("done", func() {

			_, stringIndex := taskIndex.GetCurrentOption()
			index, _ := strconv.Atoi(stringIndex)
			editTable(index, taskTitle.GetText(), taskDeadline.GetText(), taskNotes.GetText())
			renderMessage("Task edited successfully!")

		}).
		AddButton("back", func() {
			if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
				panic(err)
			}
		}).
		AddButton("quit", func() {
			renderQuit()
		})

	form.SetBorder(false).SetTitle(" edit task ").SetTitleAlign(tview.AlignCenter)

	centeredForm := Center(63, 12, form)

	if err := app.SetRoot(centeredForm, true).EnableMouse(true).SetFocus(centeredForm).Run(); err != nil {
		panic(err)
	}

}

func renderToggle() {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	taskIndexes := []string{}
	for index, _ := range todosall {
		taskIndexes = append(taskIndexes, strconv.Itoa(index))
	}

	if len(taskIndexes) == 0 {
		renderMessage("Error: No tasks to toggle, please add a task!")
	}

	taskIndex := tview.NewDropDown().
		SetLabel("Select a task to toggle completion (hit enter): ").
		SetOptions(taskIndexes, nil)

	form = tview.NewForm().
		AddFormItem(taskIndex).
		AddButton("toggle", func() {
			_, option := taskIndex.GetCurrentOption()
			indexToToggle, _ := strconv.Atoi(option)
			toggleTask(indexToToggle)
			renderMessage("Completion toggled successfully!")
		}).
		AddButton("back", func() {
			if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
				panic(err)
			}
		}).
		AddButton("quit", func() {
			renderQuit()
		})

	centeredForm := Center(55, 6, form)

	if err := app.SetRoot(centeredForm, true).EnableMouse(true).SetFocus(form).Run(); err != nil {
		panic(err)
	}

}

func renderDel() {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	if len(todosall) == 0 {
		renderMessage("Error: Nothing to delete, please add a task")
	}

	taskIndexes := []string{}
	for index, _ := range todosall {
		taskIndexes = append(taskIndexes, strconv.Itoa(index))
	}

	taskIndex := tview.NewDropDown().
		SetLabel("Select an index (hit enter): ").
		SetOptions(taskIndexes, nil)

	form = tview.NewForm().
		AddFormItem(taskIndex).
		AddButton("delete", func() {
			_, option := taskIndex.GetCurrentOption()
			indexToDel, _ := strconv.Atoi(option)
			delFromTable(indexToDel)
			renderMessage("Task deleted successfully!")
		}).
		AddButton("back", func() {
			if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
				panic(err)
			}
		}).
		AddButton("quit", func() {
			renderQuit()
		})

	centeredForm := Center(55, 6, form)

	if err := app.SetRoot(centeredForm, true).EnableMouse(true).SetFocus(form).Run(); err != nil {
		panic(err)
	}

}

func renderDelall() {

	modal := tview.NewModal().
		SetText("Delete all items in todo list").
		AddButtons([]string{"delete", "cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "delete" {
				delallFromTable()
				renderMessage("Tasks deleted successfully!")
			} else if buttonLabel == "cancel" {
				if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
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
					if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
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

func editTable(index int, title string, deadline string, notes string) {

	if title == "" {

		modal := tview.NewModal().
			SetText("title cannot be empty").
			AddButtons([]string{"ok"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "ok" {
					if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
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

	todosall[index].Title = title
	todosall[index].Deadline = deadline
	todosall[index].Notes = notes

	storage.Save(todosall)

}

func toggleTask(index int) {

	storage := storage.NewStorage[Todos]("todos.json")
	todosall := Todos{}
	storage.Load(&todosall)

	isCompleted := (todosall)[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		todosall[index].CompletedAt = &completionTime
	}

	todosall[index].Completed = !isCompleted

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
		SetText("Do you want to exit tui?").
		AddButtons([]string{"cancle", "quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "quit" {
				app.Stop()
				os.Exit(0)
			} else if buttonLabel == "cancle" {
				if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
					panic(err)
				}
			}
		})

	if err := app.SetRoot(modal, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func renderMessage(message string) {

	modal := tview.NewModal().
		SetText(message).
		AddButtons([]string{"ok"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "ok" {
				if err := app.SetRoot(centeredRoot, true).EnableMouse(true).SetFocus(centeredRoot).Run(); err != nil {
					panic(err)
				}
			}
		})

	if err := app.SetRoot(modal, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
