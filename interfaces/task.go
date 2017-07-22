package interfaces

import (
	"time"
)

type Task interface {
	GetName() string
	GetAssignee() Person
	GetDeadline() time.Time
}
