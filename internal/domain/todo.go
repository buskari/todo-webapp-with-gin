package domain

type Todo struct {
	id        int
	Title     string
	completed bool
	deleted   bool
}

func NewTodo(id int, title string) *Todo {
	return &Todo{
		id:        id,
		Title:     title,
		completed: false,
		deleted:   false,
	}
}

func (t *Todo) ID() int {
	return t.id
}

func (t *Todo) ToggleCompleted() {
	t.completed = !t.completed
}

func (t *Todo) IsCompleted() bool {
	return t.completed
}

func (t *Todo) IsDeleted() bool {
	return t.deleted
}

func (t *Todo) ToggleDeleted() {
	t.deleted = !t.deleted
}
