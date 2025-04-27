package domain

var lastID int = 0

type Todo struct {
	id        int
	Title     string
	completed bool
}

func NewTodo(title string) *Todo {
	return &Todo{
		id:        generateID(),
		Title:     title,
		completed: false,
	}
}

func (t *Todo) ID() int {
	return t.id
}

func (t *Todo) MarkAsCompleted() {
	t.completed = true
}

func (t *Todo) ToggleCompleted() {
	t.completed = !t.completed
}

func (t *Todo) IsCompleted() bool {
	return t.completed
}

func generateID() int {
	lastID++
	return lastID
}
