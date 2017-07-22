package interfaces

type Plan interface {
	Init([]Person, []Task)
	MarkTaskAsDone(Task) error
}