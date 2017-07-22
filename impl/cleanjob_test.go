package impl

import "testing"
import "time"
import "strings"

func TestNewCleaningJob(t *testing.T) {
	now := time.Now()
	name := "Hello"
	cj := NewCleanJob(name, now)
	if !strings.EqualFold(name, cj.Roomname) {
		t.Errorf("Expected %v, but Name was %v", name, cj.Roomname)
	}
	if !now.Equal(cj.LastDone) {
		t.Errorf("Expected %v, but LastDone was %v", now, cj.Roomname)
	}
}

func createCleanJobWithDate(name string, ti time.Time) *Cleanjob {
	return NewCleanJob(name, ti)
}

func getSimpleDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func TestSpecificDate(t *testing.T) {
	date := getSimpleDate(2016, 10, 13)
	cj := createCleanJobWithDate("Hello", date)
	if !date.Equal(cj.LastDone) {
		t.Errorf("Expected %v but LastDone was %v", date, cj.LastDone)
	}
}

func TestMarkAsDone(t *testing.T) {
	lastDone := getSimpleDate(2016, 10, 1)
	cj := createCleanJobWithDate("Hello", lastDone)
	cj.MarkAsDone()
	now := time.Now()
	if lastDone.Equal(cj.LastDone) {
		t.Errorf("Task was marked as done, but date was not updated.")
	}
	if now.Year() != cj.LastDone.Year() {
		t.Errorf("Expected %v but LastDone.Year was %v", now.Year(), cj.LastDone.Year())
	}
	if now.Month() != cj.LastDone.Month() {
		t.Errorf("Expected %v but LastDone.Month was %v", now.Month(), cj.LastDone.Month())
	}
	if now.Day() != cj.LastDone.Day() {
		t.Errorf("Expected %v but LastDone.Day was %v", now.Day(), cj.LastDone.Day())
	}
}

func TestOverdue(t *testing.T) {
	now := time.Now()
	lastDone := getSimpleDate(now.Year(), now.Month(), now.Day())
	cj := createCleanJobWithDate("hello", lastDone)
	if cj.getColorDependingOnTimePassed() != InWeek {
		t.Errorf("Task was overdue: %v, colored %v", cj.LastDone, cj.getColorDependingOnTimePassed())
	}
	lastYear := time.Now().AddDate(-1, 0, 0)
	lastDone = getSimpleDate(lastYear.Year(), lastYear.Month(), lastYear.Day())
	cj = createCleanJobWithDate("hello2", lastDone)
	if cj.getColorDependingOnTimePassed() != Overdue {
		t.Errorf("Task was not overdue: %v, colored %v", cj.LastDone, cj.getColorDependingOnTimePassed())
	}
	lastWeek := time.Now().AddDate(0, 0, -8)
	lastDone = getSimpleDate(lastWeek.Year(), lastWeek.Month(), lastWeek.Day())
	cj = createCleanJobWithDate("hello3", lastDone)
	if cj.getColorDependingOnTimePassed() != InLastTwoWeeks {
		t.Errorf("Task was not in lastTwoWeeks: %v, colored %v", cj.LastDone, cj.getColorDependingOnTimePassed())
	}
}
