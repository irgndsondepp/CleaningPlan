package cleaningplan

import "testing"
import "github.com/irgndsondepp/cleaningplan/people"
import "strings"

func TestRotateJob(t *testing.T) {
	mock1, tasks1 := createMockPerson("Dude")
	mock2, _ := createMockPerson("Dudette")
	cp := NewCleaningPlan()
	cp.People = []people.People{mock1, mock2}
	taskName := tasks1[len(tasks1)/2]
	err := cp.MarkJobAsDone(mock1.Name(), taskName)
	if err != nil {
		t.Errorf("Unexpected error marking job as done: %v", err)
	}
	if len(mock2.Jobs) == len(mock1.Jobs) {
		t.Errorf("Expected Mock2 to have the job now. 1: %v, 2: %v", mock1, mock2)
	}
	found := false
	for _, j := range mock2.Jobs {
		if strings.EqualFold(j.Name(), taskName) {
			found = true
		}
	}
	if !found {
		t.Errorf("Expected job to have moved to Mock2: %v", mock2.Jobs)
	}
}

func createMockPerson(name string) (*people.Mock, []string) {
	return people.NewMock(name, 5)
}
