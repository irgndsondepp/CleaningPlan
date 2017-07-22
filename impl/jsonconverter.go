package impl

import (
	"encoding/json"

	"fmt"

	"github.com/irgndsondepp/cleaningplan/interfaces"
)

type JSONConverter struct {
	taskType   interfaces.Task
	personType interfaces.Person
}

func NewJSONConverter() *JSONConverter {
	return &JSONConverter{}
}

func (*JSONConverter) ConvertTo(plan interfaces.Plan) ([]byte, error) {
	return json.MarshalIndent(plan, "", "\t")
}

func (*JSONConverter) ReadFrom(bytes []byte, cp interfaces.Plan) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(bytes, &objMap)
	if err != nil {
		return err
	}
	tasks, err := unmarshalTasks(objMap)
	if err != nil {
		return fmt.Errorf("error parsing tasks: %v", err)
	}

	people, err := unmarshalPersons(objMap)
	if err != nil {
		return fmt.Errorf("error parsing persons: %v", err)
	}

	cp.Init(people, tasks)
	return err
}

func unmarshalTasks(objMap map[string]*json.RawMessage) ([]interfaces.Task, error) {
	var taskMap []*json.RawMessage
	err := json.Unmarshal(*objMap["tasks"], &taskMap)
	if err != nil {
		return nil, err
	}
	var tasks []interfaces.Task
	for _, item := range taskMap {
		cj, err := unmarshalSingleTask(item)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, cj)
	}
	return tasks, nil
}

func unmarshalSingleTask(item *json.RawMessage) (interfaces.Task, error) {
	cj := &Cleanjob{}
	err := json.Unmarshal(*item, &cj)
	if err != nil {
		return nil, err
	}
	return cj, nil
}

func unmarshalPersons(objMap map[string]*json.RawMessage) ([]interfaces.Person, error) {
	var peopleMap []*json.RawMessage
	err := json.Unmarshal(*objMap["people"], &peopleMap)
	if err != nil {
		return nil, err
	}
	var people []interfaces.Person
	for _, item := range peopleMap {
		p, err := unmarshalSinglePerson(item)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}

func unmarshalSinglePerson(item *json.RawMessage) (interfaces.Person, error) {
	p := &Flatmate{}
	err := json.Unmarshal(*item, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
