package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/buskari/todo-webapp-with-gin/internal/domain"
)

var (
	todoDB []*domain.Todo
	once   sync.Once
)

func BuildDB() []*domain.Todo {
	once.Do(func() {
		file, err := os.Open("internal/data/todos_seed.csv")
		if err != nil {
			log.Fatalf("Could not read from source file: %v", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		tasks, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		todoDB = []*domain.Todo{}
		for id, task := range tasks {
			id += 1
			todo := domain.NewTodo(id, task[1])
			todoDB = append(todoDB, todo)
		}

		fmt.Printf("Loaded %d ToDos from CSV\n", len(todoDB))
	})

	return todoDB
}
