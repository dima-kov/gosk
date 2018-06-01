package tasks

import "time"

type Broker interface {
	AddTask(task Task, delay time.Duration, args ...interface{})
	HandleWaitingQueue()
	serializeTask(task Task, uuid string, args ...interface{}) ([]byte, error)
}
