package cleaningplan

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
	"time"
)

type CleaningPlan struct {
	Flatmates []*Flatmate `xml:"Flatmates" json:"Flatmates"`
}

func NewCleaningPlan() *CleaningPlan {
	cp := CleaningPlan{}
	return &cp
}

func InitCleaningPlan() *CleaningPlan {
	benni := NewFlatmate("Benni")
	benni.AddCleanJob(NewCleanJob("LivingRoom", SimpleDate(2017, 3, 21)))

	markus := NewFlatmate("Markus")
	markus.AddCleanJob(NewCleanJob("Kitchen", SimpleDate(2016, 10, 3)))
	markus.AddCleanJob(NewCleanJob("Bath", SimpleDate(2017, 5, 27)))

	robert := NewFlatmate("Robert")

	fms := []*Flatmate{markus, benni, robert}
	cp := CleaningPlan{
		Flatmates: fms,
	}
	return &cp
}

func (cp *CleaningPlan) MarkJobAsDone(flatMate, roomName string) error {
	var job *Cleanjob
	var err error
	j := -1
	for i, fm := range cp.Flatmates {
		if strings.ToLower(fm.Name) == strings.ToLower(flatMate) {
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
	j = (j + 1) % len(cp.Flatmates)
	cp.Flatmates[j].AddCleanJob(job)
	return nil
}

func SimpleDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 12, 0, 0, 0, time.Local)
}

func (cp *CleaningPlan) ToString() string {
	msg := ""
	for _, fm := range cp.Flatmates {
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
	cleaningPlan := CleaningPlan{}
	err := json.Unmarshal(bytes, &cleaningPlan)
	return &cleaningPlan, err
}

func FromXML(bytes []byte) (*CleaningPlan, error) {
	cleaningPlan := CleaningPlan{}
	err := xml.Unmarshal(bytes, &cleaningPlan)
	return &cleaningPlan, err
}
