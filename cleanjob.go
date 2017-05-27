package cleaningplan

import "time"
import "fmt"

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

func (c *Cleanjob) ToString() string {
	return fmt.Sprintf("%v (last done on %v)", c.Roomname, c.LastDone)
}
