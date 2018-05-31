package tasks

import (
	"github.com/pkg/errors"
	"time"
)

const TaskNameStartsWith = "tasks_"

type TaskManager interface {
	Delay(task Task, delay time.Duration, args ...interface{})
}

type taskManager struct {
	broker Broker
}

func NewTaskManager(brokerType, host string, port uint) (TaskManager, error) {
	var broker Broker
	switch brokerType {
	case "redis":
		broker = NewRedisBroker(host, port, "")
	default:
		return nil, errors.New("undefined broker type: " + brokerType)
	}
	return taskManager{broker: broker}, nil
}

func (tm taskManager) Delay(task Task, delay time.Duration, args ...interface{}) {
	tm.broker.SetTask(task, delay, args)
}
