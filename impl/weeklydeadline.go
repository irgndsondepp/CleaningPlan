package impl

import (
	"time"
)

type WeeklyDeadline struct {
	LastDone time.Time
}

func (w *WeeklyDeadline) IsOverdue() bool {
	if time.Now().Sub(w.LastDone) > time.Hour*24*7 {
		return true
	}
	return false
}

func (w *WeeklyDeadline) IsLongOverdue() bool {
	if time.Now().Sub(w.LastDone) > time.Hour*24*14 {
		return true
	}
	return false
}

func (w *WeeklyDeadline) GetDeadline() time.Time {
	return w.LastDone
}
