package impl

import "time"

import "github.com/irgndsondepp/cleaningplan/interfaces"

type Cleanjob struct {
	Roomname string            `json:"RoomName"`
	Deadline time.Time         `json:"LastDone"`
	Assignee interfaces.Person `json:"Assignee"`
}

func NewCleanjob(roomName string, lastDone time.Time, assignee interfaces.Person) *Cleanjob {
	cj := Cleanjob{
		Roomname: roomName,
		Deadline: lastDone,
		Assignee: assignee,
	}
	return &cj
}

func (cj *Cleanjob) GetName() string {
	return cj.Roomname
}

func (cj *Cleanjob) GetAssignee() interfaces.Person {
	return cj.Assignee
}

func (cj *Cleanjob) GetDeadline() time.Time {
	return cj.Deadline
}
