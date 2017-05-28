package cleaningplan

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
	"time"

	"github.com/irgndsondepp/cleaningplan/people"
	"github.com/irgndsondepp/cleaningplan/people/tasks"
)

type CleaningPlan struct {
	People []people.People `xml:"People" json:"People"`
}

type specificCleaningPlan struct {
	People []*people.SpecificFlatmate `xml:"People" json:"People"`
}

func newSpecificCleaningPlan() *specificCleaningPlan {
	return &specificCleaningPlan{}
}

func (s *specificCleaningPlan) ToSimpleCleaningPlan() *CleaningPlan {
	cp := CleaningPlan{}
	var pep []people.People
	for _, p := range s.People {
		pep = append(pep, p.ToSimpleFlatmate())
	}
	cp.People = pep
	return &cp
}

func NewCleaningPlan() *CleaningPlan {
	cp := CleaningPlan{}
	return &cp
}

func InitCleaningPlan() *CleaningPlan {
	benni := people.NewFlatmate("Benni")
	benni.AddJob(tasks.NewCleanJob("LivingRoom", SimpleDate(2017, 3, 21)))

	markus := people.NewFlatmate("Markus")
	markus.AddJob(tasks.NewCleanJob("Kitchen", SimpleDate(2016, 10, 3)))
	markus.AddJob(tasks.NewCleanJob("Bath", SimpleDate(2017, 5, 27)))

	robert := people.NewFlatmate("Robert")

	fms := []people.People{markus, benni, robert}
	cp := CleaningPlan{
		People: fms,
	}
	return &cp
}

func (cp *CleaningPlan) MarkJobAsDone(flatMate, roomName string) error {
	var job tasks.Doable
	var err error
	j := -1
	for i, fm := range cp.People {
		if strings.ToLower(fm.Name()) == strings.ToLower(flatMate) {
			job, err = fm.MarkJobAsDone(roomName)
			if err != nil {
				return err
			}
			j = i
		}
	}
	if job == nil {
		return errors.New("No flatmate found.")
	}
	j = (j + 1) % len(cp.People)
	cp.People[j].AddJob(job)
	return nil
}

func SimpleDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 12, 0, 0, 0, time.Local)
}

func (cp *CleaningPlan) ToString() string {
	msg := ""
	for _, fm := range cp.People {
		msg += fm.ToString()
	}
	return msg
}

func (cp *CleaningPlan) ToXML() ([]byte, error) {
	return xml.MarshalIndent(cp, "", "\t")
}

func (cp *CleaningPlan) ToJSON() ([]byte, error) {
	return json.MarshalIndent(cp, "", "\t")
}

func FromJSON(bytes []byte) (*CleaningPlan, error) {
	cleaningPlan := newSpecificCleaningPlan()
	err := json.Unmarshal(bytes, &cleaningPlan)
	return cleaningPlan.ToSimpleCleaningPlan(), err
}

func FromXML(bytes []byte) (*CleaningPlan, error) {
	cleaningPlan := newSpecificCleaningPlan()
	err := xml.Unmarshal(bytes, &cleaningPlan)
	return cleaningPlan.ToSimpleCleaningPlan(), err
}
