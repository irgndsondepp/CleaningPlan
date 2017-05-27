package cleaningplan

import (
	"errors"
	"fmt"
	"strings"
)

type Flatmate struct {
	Name string      `xml:"Name" json:"Name"`
	Jobs []*Cleanjob `xml:"Cleanjobs" json:"Cleanjobs"`
}

func NewFlatmate(name string) *Flatmate {
	//cleanJobs := make([]*Cleanjob, 1)
	fm := Flatmate{
		Name: name,
	}
	return &fm
}

func (f *Flatmate) AddCleanJob(cj *Cleanjob) {
	f.Jobs = append(f.Jobs, cj)
}

func (f *Flatmate) MarkJobAsDone(roomName string) (*Cleanjob, error) {
	var doneJob *Cleanjob
	var cleanJobs []*Cleanjob
	for _, job := range f.Jobs {
		if strings.ToLower(job.Roomname) != strings.ToLower(roomName) {
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

func (f *Flatmate) ToString() string {
	if len(f.Jobs) < 1 {
		return fmt.Sprintf("Nothing to do for %v\n", f.Name)
	}
	msg := ""
	for _, j := range f.Jobs {
		msg += fmt.Sprintf("%v is up for cleaning the %v\n", f.Name, j.ToString())
	}
	return msg
}
