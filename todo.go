package main

import( 

    "github.com/hardikkum444/go-do-it/cmd"
    "time"
    
)

type Todo struct{
    Title string
    Completed bool
    CreatedAt time.Time
    CompletedAt *time.Time
}

type Todos []Todo 

var todos Todos

func(todos *Todos) add (title string) {

    todo := Todo{
        Title : title,
        Completed : false,
        CreatedAt : time.Now().UTC(),
        CompletedAt : nil,
    }

    *todos = append(*todos, todo)
}



