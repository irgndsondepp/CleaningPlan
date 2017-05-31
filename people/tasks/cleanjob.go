package tasks

import "time"
import "fmt"

type Doable interface {
	Name() string
	MarkAsDone()
	ToString() string
	ToHtml() string
}

type Cleanjob struct {
	Roomname string    `xml:"RoomName" json:"RoomName"`
	LastDone time.Time `xml:"LastDone" json:"LastDone"`
}

func NewCleanJob(roomName string, lastDone time.Time) *Cleanjob {
	cj := Cleanjob{
		Roomname: roomName,
		LastDone: lastDone,
	}
	return &cj
}

func (cj *Cleanjob) MarkAsDone() {
	cj.LastDone = time.Now()
}

func (cj *Cleanjob) Name() string {
	return cj.Roomname
}

func (c *Cleanjob) ToString() string {
	return fmt.Sprintf("%v (last done on %v)", c.Roomname, c.LastDone)
}

func (c *Cleanjob) ToHtml() string {
	return fmt.Sprintf("<p>%v (last done on %v)</p>", c.Roomname, c.LastDone)
}
