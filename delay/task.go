package delay

import "time"

type DelayTask interface {
	Delay(t time.Duration, args ...interface{})
	RunTask(args ...interface{})
	GetName() string
}
