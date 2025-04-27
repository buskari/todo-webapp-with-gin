package persistence

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

type InMemoryRepository struct {
	todos []*domain.Todo
}

func buildDB() []*domain.Todo {
	once.Do(func() {
		file, err := os.Open("data/todos_seed.csv")
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
		for _, task := range tasks {
			todo := domain.NewTodo(task[0])
			todoDB = append(todoDB, todo)
		}

		fmt.Printf("Loaded %d ToDos from CSV\n", len(todoDB))
	})

	return todoDB
}

func NewTodoRepository() *InMemoryRepository {
	return &InMemoryRepository{
		todos: buildDB(),
	}
}
