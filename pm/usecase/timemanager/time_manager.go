package timemanager

import "time"

type TimeManager interface {
	Now() time.Time
}
