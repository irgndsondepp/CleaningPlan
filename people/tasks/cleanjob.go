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
	return fmt.Sprintf("%v (last done on %v)", c.Roomname, c.LastDone.Format(time.UnixDate))
}

func (c *Cleanjob) ToHtml() string {
	return fmt.Sprintf("<p>%v</p>", c.colorLastDoneInHtml())
}

func (c *Cleanjob) colorLastDoneInHtml() string {
	if c.isOverdue() {
		return fmt.Sprintf("<font color=\"red\">%v</font>", c.ToString())
	} else {
		return fmt.Sprintf("<font color=\"green\">%v</font>", c.ToString())
	}
}

func (c *Cleanjob) isOverdue() bool {
	aWeekAgo := time.Now().AddDate(0, 0, -7)
	if aWeekAgo.Before(c.LastDone) {
		return false
	}
	return true
}
