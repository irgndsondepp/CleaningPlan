package interfaces

import (
	"time"
)

type Task interface {
	GetName() string
	GetAssignee() string
	GetDeadline() time.Time
}
