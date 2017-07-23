package interfaces

type Plan interface {
	Init([]Person, []Task)
	MarkTaskAsDone(Task) error
	GetTasks() []Task
	FilterTasks(string) (interface{}, error)
}
