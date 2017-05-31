package people

import (
	"errors"
	"fmt"
	"strings"

	"github.com/irgndsondepp/cleaningplan/people/tasks"
)

type People interface {
	Name() string
	AddJob(job tasks.Doable)
	MarkJobAsDone(jobName string) (tasks.Doable, error)
	ToString() string
	ToHtml() string
}

type Flatmate struct {
	PersonName string         `xml:"Name" json:"Name"`
	Jobs       []tasks.Doable `xml:"Cleanjobs" json:"Cleanjobs"`
}

type SpecificFlatmate struct {
	PersonName string            `xml:"Name" json:"Name"`
	Jobs       []*tasks.Cleanjob `xml:"Cleanjobs" json:"Cleanjobs"`
}

func (s *SpecificFlatmate) ToSimpleFlatmate() *Flatmate {
	var jobs []tasks.Doable
	for _, j := range s.Jobs {
		jobs = append(jobs, j)
	}
	fm := Flatmate{
		PersonName: s.PersonName,
		Jobs:       jobs,
	}
	return &fm
}

func NewFlatmate(name string) *Flatmate {
	fm := Flatmate{
		PersonName: name,
	}
	return &fm
}

func (f *Flatmate) AddJob(cj tasks.Doable) {
	f.Jobs = append(f.Jobs, cj)
}

func (f *Flatmate) MarkJobAsDone(taskName string) (tasks.Doable, error) {
	var doneJob tasks.Doable
	var cleanJobs []tasks.Doable
	for _, job := range f.Jobs {
		if strings.ToLower(job.Name()) != strings.ToLower(taskName) {
			cleanJobs = append(cleanJobs, job)
		} else {
			job.MarkAsDone()
			doneJob = job
		}
	}
	if doneJob != nil {
		f.Jobs = cleanJobs
		return doneJob, nil
	}
	return nil, errors.New("No job found.")
}

func (f *Flatmate) Name() string {
	return f.PersonName
}

func (f *Flatmate) ToString() string {
	if len(f.Jobs) < 1 {
		return fmt.Sprintf("Nothing to do for %v\n", f.PersonName)
	}
	msg := ""
	for _, j := range f.Jobs {
		msg += fmt.Sprintf("%v is up for cleaning the %v\n", f.PersonName, j.ToString())
	}
	return msg
}

func (f *Flatmate) ToHtml() string {
	if len(f.Jobs) < 1 {
		return fmt.Sprintf("<p>Nothing to do for %v</p>", f.PersonName)
	}
	msg := ""
	for _, j := range f.Jobs {
		msg += fmt.Sprintf("<item>%v</item>", j.ToHtml())
	}
	msg = fmt.Sprintf("<ul>%v</ul>", msg)
	return fmt.Sprintf("<p>%v is up for the following tasks:</p><p>%v</p>", f.PersonName, msg)
}
