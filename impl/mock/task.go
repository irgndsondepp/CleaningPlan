package mock

import "time"
import "fmt"

var i = 0

type Task struct {
	Name     string
	Assignee string
}

func NewMockTask(assignee string) *Task {
	t := &Task{
		Name:     fmt.Sprintf("Task Nr. %v", i),
		Assignee: assignee,
	}
	i++
	return t
}

func (t *Task) GetName() string {
	return t.Name
}
func (t *Task) GetAssignee() string {
	return t.Assignee
}
func (*Task) GetDeadline() time.Time {
	return time.Now()
}
