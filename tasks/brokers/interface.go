package brokers

import (
	"github.com/dima-kov/go-tasks/tasks"
	"time"
)

type Broker interface {
	AddTask(task tasks.Task, delay time.Duration, args ...interface{}) (int64, error)
	HandleWaitingQueue()
	serializeTask(task tasks.Task, uuid string, args ...interface{}) ([]byte, error)
}
