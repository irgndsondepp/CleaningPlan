package impl

import "github.com/irgndsondepp/cleaningplan/interfaces"
import "time"

type SimpleTask struct {
	Name     string
	Assignee interfaces.Person
}

func NewSimpleTask(taskName, assignee string) *SimpleTask {
	return &SimpleTask{
		Name:     taskName,
		Assignee: NewSimplePerson(assignee),
	}
}

func (s *SimpleTask) GetName() string {
	return s.Name
}

func (s *SimpleTask) GetAssignee() interfaces.Person {
	return s.Assignee
}

func (*SimpleTask) GetDeadline() time.Time {
	return time.Now()
}
