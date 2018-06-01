package task

import "time"

type Task interface {
	Delay(t time.Duration, args ...interface{})
	RunTask(args ...interface{})
	GetName() string
}
