package people

import (
	"errors"
	"fmt"
	"strings"

	"github.com/irgndsondepp/cleaningplan/people/tasks"
)

type Mock struct {
	name string
	Jobs []tasks.Doable
}

func NewMock(name string, taskCount int) (*Mock, []string) {
	var jobs []tasks.Doable
	var taskNames []string
	for i := 0; i < taskCount; i++ {
		taskName := fmt.Sprintf("%v-%v", name, string(i))
		taskNames = append(taskNames, taskName)
		jobs = append(jobs, tasks.NewMock(taskName))
	}
	m := Mock{
		name: name,
		Jobs: jobs,
	}
	return &m, taskNames
}

func (m *Mock) Name() string {
	return m.name
}

func (m *Mock) AddJob(job tasks.Doable) {
	m.Jobs = append(m.Jobs, job)
}

func (m *Mock) MarkJobAsDone(jobName string) (tasks.Doable, error) {
	var newJobs []tasks.Doable
	var job tasks.Doable
	for _, j := range m.Jobs {
		if strings.EqualFold(j.Name(), jobName) {
			job = j
		} else {
			newJobs = append(newJobs, j)
		}
	}
	if job != nil {
		m.Jobs = newJobs
		return job, nil
	}
	return nil, errors.New("No job found.")
}
func (m *Mock) ToString() string {
	return fmt.Sprintf("Name: %v, Jobs: %v", m.name, m.Jobs)
}
