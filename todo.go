package main

import( 

    "github.com/hardikkum444/go-do-it/cmd"
    "time"
    "errors"
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





