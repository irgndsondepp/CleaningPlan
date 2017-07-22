package impl

import "time"

type SimpleTask struct {
	Name     string
	Assignee string
}

func NewSimpleTask(taskName, assignee string) *SimpleTask {
	return &SimpleTask{
		Name:     taskName,
		Assignee: assignee,
	}
}

func (s *SimpleTask) GetName() string {
	return s.Name
}

func (s *SimpleTask) GetAssignee() string {
	return s.Assignee
}

func (*SimpleTask) GetDeadline() time.Time {
	return time.Now()
}
