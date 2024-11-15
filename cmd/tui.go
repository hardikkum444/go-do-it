/*
Copyright © 2024 man44 <man44@tutamail.com>

*/
package cmd

import (
	// "fmt"

	"github.com/spf13/cobra"
    "github.com/rivo/tview"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "open go-do-it tui",
    Long: "open the terminal user interface",
	Run: func(cmd *cobra.Command, args []string) {
        menu()
	},
}

var list *tview.List

func createMenuList(app *tview.Application) *tview.List {

    list = tview.NewList().
    AddItem("add a task", "", 'n', nil).
    AddItem("quit", "", 'q', func() {
        app.Stop()
    })

    return list
}

func menu() {
    
    app := tview.NewApplication()

   list = createMenuList(app) 

   if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil{
       panic(err)
   }
    
}

