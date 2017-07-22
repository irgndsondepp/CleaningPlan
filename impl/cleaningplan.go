package impl

import (
	"time"

	"fmt"

	"github.com/irgndsondepp/cleaningplan/interfaces"
)

type RotatingCleaningPlan struct {
	Tasks       []interfaces.Task   `json:"tasks"`
	People      []interfaces.Person `json:"people"`
	persistence interfaces.Persistence
}

func NewRotatingCleaningPlan(pers interfaces.Persistence) *RotatingCleaningPlan {
	cp := &RotatingCleaningPlan{}
	pers.Load(cp)
	cp.persistence = pers
	return cp
}

func (cp *RotatingCleaningPlan) Init(people []interfaces.Person, tasks []interfaces.Task) {
	cp.People = people
	cp.Tasks = tasks
}

func (cp *RotatingCleaningPlan) MarkTaskAsDone(doneTask interfaces.Task) error {
	var tasks []interfaces.Task
	var dT interfaces.Task
	for _, t := range cp.Tasks {
		if t.GetName() == doneTask.GetName() {
			dT = t
		} else {
			tasks = append(tasks, t)
		}
	}
	if dT == nil {
		return fmt.Errorf("task %v was not found", doneTask.GetName())
	}
	if dT.GetAssignee() != doneTask.GetAssignee() {
		return fmt.Errorf("task %v is not assigned to %v", dT.GetName(), doneTask.GetAssignee())
	}
	newAssignee, err := cp.getNextAssignee(dT.GetAssignee())
	if err != nil {
		return err
	}
	newTask := NewCleanjob(dT.GetName(), time.Now(), newAssignee.GetName())
	tasks = append(tasks, newTask)
	cp.Tasks = tasks
	cp.persistence.Save(cp)
	return nil
}

func (cp *RotatingCleaningPlan) FilterTasks(filter string) ([]interfaces.Task, error) {
	var tasks []interfaces.Task
	exists := false
	for _, p := range cp.People {
		if p.GetName() == filter {
			exists = true
			break
		}
	}
	if !exists {
		return nil, fmt.Errorf("Person %v not found in list", filter)
	}
	for _, t := range cp.Tasks {
		if t.GetAssignee() == filter {
			tasks = append(tasks, t)
		}
	}
	return tasks, nil
}

func (cp *RotatingCleaningPlan) getNextAssignee(lastAssignee string) (interfaces.Person, error) {
	index := 0
	for i, p := range cp.People {
		if p.GetName() == lastAssignee {
			index = i + 1
			break
		}
	}
	if index >= len(cp.People) {
		index -= len(cp.People)
	}
	for i, p := range cp.People {
		if i == index {
			return p, nil
		}
	}
	return nil, fmt.Errorf("no new assignee found")
}
