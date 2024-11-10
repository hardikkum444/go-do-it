package models

import( 

    // "github.com/hardikkum444/go-do-it/cmd"
    "time"
    "errors"
    "github.com/aquasecurity/table" 
    "os"
    "strconv"
    
)

type Todo struct{
    Title string
    Completed bool
    CreatedAt time.Time
    CompletedAt *time.Time
}

type Todos []Todo 

var todos Todos

func(todos *Todos) add(title string) {

    todo := Todo{
        Title : title,
        Completed : false,
        CreatedAt : time.Now().UTC(),
        CompletedAt : nil,
    }

    *todos = append(*todos, todo)
}

func(todos *Todos) validateIndex(index int) error {

    if index < 0 || index >= len(*todos){
        return errors.New("invalid index")
    }
    return nil
}

func(todos *Todos) delete(index int) error {

    if err := todos.validateIndex(index); err != nil {
        return err 
    }

    *todos = append((*todos)[:index], (*todos)[index+1:]...)
    return nil
}

func(todos *Todos) toggle(index int) error {

    if err := todos.validateIndex(index); err != nil {
        return err 
    }

    isCompleted := (*todos)[index].Completed
    if !isCompleted {
        completionTime := time.Now()
        (*todos)[index].CompletedAt = &completionTime
    }

    (*todos)[index].Completed = !isCompleted
    return nil
}

func(todos *Todos) edit(index int, title string) error {

    if err := todos.validateIndex(index); err != nil {
        return err 
    }

    (*todos)[index].Title = title
    return nil
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
