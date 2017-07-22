package impl

import (
	"testing"

	"github.com/irgndsondepp/cleaningplan/impl/mock"
	"github.com/irgndsondepp/cleaningplan/interfaces"
)

func TestRotateJob(t *testing.T) {
	mock1 := NewFlatmate("Dude")
	mock2 := NewFlatmate("Dudette")
	task1 := createMockTask(mock1.GetName())
	task2 := createMockTask(mock2.GetName())
	peops := []interfaces.Person{mock1, mock2}
	tasks := []interfaces.Task{task1, task2}
	cp := NewRotatingCleaningPlan(mock.NewMockPersistence(&RotatingCleaningPlan{}))
	cp.Init(peops, tasks)
	err := cp.MarkTaskAsDone(task1)
	if err != nil {
		t.Errorf("Unexpected error marking job as done: %v", err)
	}
	for _, task := range cp.Tasks {
		if task.GetName() == task1.GetName() {
			if task.GetAssignee() == task1.GetAssignee() {
				t.Errorf("Wrong assignee for %v: expected %v, but was %v", task.GetName(), mock2.GetName(), task.GetAssignee())
			} else if task.GetAssignee() != mock2.GetName() {
				t.Errorf("Expected job to have moved to Mock2: %v", task.GetAssignee())
			}
			if task.GetDeadline() == task1.GetDeadline() {
				t.Errorf("Expected new Deadline %v, but got %v", task.GetDeadline(), task1.GetDeadline())
			}
		}
	}
}

func TestRotateJobOutOfBounds(t *testing.T) {
	mock1 := NewFlatmate("Dude")
	mock2 := NewFlatmate("Dudette")
	task1 := createMockTask(mock1.GetName())
	task2 := createMockTask(mock2.GetName())
	peops := []interfaces.Person{mock1, mock2}
	tasks := []interfaces.Task{task1, task2}
	cp := NewRotatingCleaningPlan(mock.NewMockPersistence(&RotatingCleaningPlan{}))
	cp.Init(peops, tasks)
	err := cp.MarkTaskAsDone(task2)
	if err != nil {
		t.Errorf("Unexpected error marking job as done: %v", err)
	}
	for _, task := range cp.Tasks {
		if task.GetName() == task2.GetName() {
			if task.GetAssignee() == task2.GetAssignee() {
				t.Errorf("Wrong assignee for %v: expected %v, but was %v", task.GetName(), mock1.GetName(), task.GetAssignee())
			} else if task.GetAssignee() != mock1.GetName() {
				t.Errorf("Expected job to have moved to Mock1: %v", task.GetAssignee())
			}
		}
	}
}

func createMockTask(assignee string) *mock.Task {
	return mock.NewMockTask(assignee)
}
