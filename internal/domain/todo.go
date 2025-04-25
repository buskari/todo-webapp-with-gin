package domain

var lastID int = 0

type Todo struct {
	id        int
	title     string
	completed bool
}

// constructor
func NewTodo(title string) *Todo {
	return &Todo{
		id:        generateID(),
		title:     title,
		completed: false,
	}
}

// getters
func (t *Todo) ID() int {
	return t.id
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) IsCompleted() bool {
	return t.completed
}

// Setters
func (t *Todo) SetTitle(title string) {
	t.title = title
}

func (t *Todo) ToggleCompleted() {
	t.completed = !t.completed
}

// internal
func generateID() int {
	lastID++
	return lastID
}
