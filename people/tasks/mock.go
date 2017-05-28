package tasks

import (
	"fmt"
	"time"
)

type Mock struct {
	name     string
	LastDone time.Time
}

func NewMock(name string) *Mock {
	m := Mock{
		name:     name,
		LastDone: time.Now(),
	}
	return &m
}

func (m *Mock) Name() string {
	return m.name
}

func (m *Mock) MarkAsDone() {
	m.LastDone = time.Now()
}

func (m *Mock) ToString() string {
	return fmt.Sprintf("Name: %v, LastDone: %v", m.name, m.LastDone)
}
