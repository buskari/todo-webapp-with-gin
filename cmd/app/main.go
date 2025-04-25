package main

import (
	"fmt"

	"github.com/buskari/todo-webapp-with-gin/internal/infra/persistence"
)

func main() {
	todos := persistence.BuildDB()
	for _, todo := range todos {
		if todo.ID() == 3 {
			todo.ToggleCompleted()
		}
	}

	fmt.Println(todos[2].IsCompleted())
}
