package tasks

import "time"

type Broker interface {
	SetTask(task Task, delay time.Duration, args ...interface{})
	serializeTask(task Task, args ...interface{}) ([]byte, error)
}
