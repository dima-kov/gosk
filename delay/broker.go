package delay

import (
	"time"
)

type Broker interface {
	AddTask(task DelayTask, delay time.Duration, args ...interface{}) (int64, error)
	HandleWaitingQueue()
	serializeTask(task DelayTask, uuid string, args ...interface{}) ([]byte, error)
}
