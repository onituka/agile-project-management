package clock

import "time"

type realTime struct{}

func NewRealTime() *realTime {
	return &realTime{}
}

func (c *realTime) Now() time.Time {
	return time.Now().UTC()
}
