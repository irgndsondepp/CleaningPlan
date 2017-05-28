package people

import (
	"strings"
	"testing"

	"github.com/irgndsondepp/cleaningplan/people/tasks"
)

var taskName = "Hello"

func TestNewFlatmate(t *testing.T) {
	name := "Dude"
	fm := NewFlatmate(name)
	if !strings.EqualFold(name, fm.Name) {
		t.Errorf("Expected %v but Name was %v", name, fm.Name)
	}
	if len(fm.Jobs) != 0 {
		t.Errorf("Jobs should be empty, but were %v", fm.Jobs)
	}
}

func TestAddJob(t *testing.T) {
	fm := createFlatmateWithTask()
	if len(fm.Jobs) != 1 {
		t.Errorf("1 Job was added, but List contains %v elements: %v", len(fm.Jobs), fm.Jobs)
	}
}

func createFlatmateWithTask() *Flatmate {
	fm := NewFlatmate("Dude")
	fm.AddJob(tasks.NewMock(taskName))
	return fm
}

func TestDoJob(t *testing.T) {
	fm := createFlatmateWithTask()
	res, err := fm.MarkJobAsDone(taskName)
	if err != nil {
		t.Errorf("Unexpected error while marking job as done: %v", err)
	}
	if res == nil {
		t.Errorf("Result was empty, but no Error occured")
	}
}

func TestDoOneOfManyJobs(t *testing.T) {
	fm := createFlatmateWithTask()
	jobNames := []string{"Job2", "Job3", "Job4"}
	for _, name := range jobNames {
		fm.AddJob(tasks.NewMock(name))
	}
	jobToDo := jobNames[1]
	res, err := fm.MarkJobAsDone(jobToDo)
	if err != nil {
		t.Errorf("Unexpected error while marking job as done: %v", err)
	}
	if res == nil {
		t.Errorf("Result was empty, but no Error occured")
	}
	if !strings.EqualFold(jobToDo, res.Name()) {
		t.Errorf("Wrong job done. Expected %v, but was %v", jobToDo, res.Name())
	}
}

func TestRemoveJobTwice(t *testing.T) {
	fm := createFlatmateWithTask()
	_, err := fm.MarkJobAsDone(taskName)
	if err != nil {
		t.Errorf("Unexpected error during first job completion: %v", err)
	}
	res, err := fm.MarkJobAsDone(taskName)
	if res != nil {
		t.Errorf("Result should be empty, but was %v", res)
	}
	if err == nil {
		t.Errorf("No Error occured during second completion of done task.")
	}
}
