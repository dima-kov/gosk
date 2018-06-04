package delay

import (
	"time"
)

type Broker interface {
	AddTask(task Task, delay time.Duration, args ...interface{}) (int64, error)
	HandleWaitingQueue()
	serializeTask(task Task, uuid string, args ...interface{}) ([]byte, error)
}
