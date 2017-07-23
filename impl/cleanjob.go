package impl

import "time"

type Cleanjob struct {
	Roomname string    `json:"roomname"`
	Deadline time.Time `json:"deadline"`
	Assignee string    `json:"assignee"`
}

func NewCleanjob(roomName string, lastDone time.Time, assignee string) *Cleanjob {
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

func (cj *Cleanjob) GetAssignee() string {
	return cj.Assignee
}

func (cj *Cleanjob) GetDeadline() time.Time {
	return cj.Deadline
}
