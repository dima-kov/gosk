package delay

import (
	"github.com/pkg/errors"
	"time"
)

const WaitingQueueName = "go_tasks"

type TaskManager interface {
	Delay(task Task, delay time.Duration, args ...interface{})
	RegisterTasks(tasks ...Task)
	HandleWaitingQueue()
}

type taskManager struct {
	broker          Broker
	registeredTasks map[string]Task
}

func NewTaskManager(brokerType, host string, port uint) (TaskManager, error) {
	var broker Broker
	switch brokerType {
	case "redis":
		broker = NewRedisBroker(host, port, "", "")
	default:
		return nil, errors.New("undefined broker type: " + brokerType)
	}
	manager := taskManager{broker: broker, registeredTasks: map[string]Task{}}
	manager.HandleWaitingQueue()
	return &manager, nil
}

func (tm *taskManager) RegisterTasks(tasks ...Task) {
	for _, item := range tasks {
		tm.registeredTasks[item.GetName()] = item
	}
}

func (tm *taskManager) Delay(task Task, delay time.Duration, args ...interface{}) {
	tm.broker.AddTask(task, delay, args)
}

func (tm *taskManager) HandleWaitingQueue() {
	tm.broker.HandleWaitingQueue()
}
