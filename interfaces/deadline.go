package interfaces

import (
	"time"
)

type DeadlineDetector interface {
	IsOverdue() bool
	IsLongOverdue() bool
	GetDeadline() *time.Time
}
