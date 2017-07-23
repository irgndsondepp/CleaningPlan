package impl

import (
	"time"

	"fmt"

	"strings"

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
	dtn := strings.ToLower(doneTask.GetName())
	for _, t := range cp.Tasks {
		if strings.ToLower(t.GetName()) == dtn {
			dT = t
		} else {
			tasks = append(tasks, t)
		}
	}
	if dT == nil {
		return fmt.Errorf("task %v was not found", doneTask.GetName())
	}
	dta := strings.ToLower(doneTask.GetAssignee())
	if strings.ToLower(dT.GetAssignee()) != dta {
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

func (cp *RotatingCleaningPlan) FilterTasks(filter string) (interface{}, error) {
	tasks := make([]*FilteredTask, 0)
	exists := false
	f := strings.ToLower(filter)
	for _, p := range cp.People {
		if strings.ToLower(p.GetName()) == f {
			exists = true
			break
		}
	}
	if !exists {
		return nil, fmt.Errorf("Person %v not found in list", filter)
	}
	for _, t := range cp.Tasks {
		if strings.ToLower(t.GetAssignee()) == f {
			ft := &FilteredTask{Name: t.GetName(), Deadline: t.GetDeadline()}
			tasks = append(tasks, ft)
		}
	}
	return tasks, nil
}

func (cp *RotatingCleaningPlan) GetTasks() []interfaces.Task {
	return cp.Tasks
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

type FilteredTask struct {
	Name     string    `json:"roomname"`
	Deadline time.Time `json:"deadline"`
}
