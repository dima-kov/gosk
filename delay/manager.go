package delay

import (
	"github.com/pkg/errors"
	"time"
)

const WaitingQueueName = "go_tasks"

type TaskManager interface {
	Delay(task DelayTask, delay time.Duration, args ...interface{})
	RegisterTasks(tasks ...DelayTask)
	HandleWaitingQueue()
}

type taskManager struct {
	broker          Broker
	registeredTasks map[string]DelayTask
}

func NewTaskManager(brokerType, host string, port uint) (TaskManager, error) {
	var broker Broker
	switch brokerType {
	case "redis":
		broker = NewRedisBroker(host, port, "", "")
	default:
		return nil, errors.New("undefined broker type: " + brokerType)
	}
	manager := taskManager{broker: broker, registeredTasks: map[string]DelayTask{}}
	manager.HandleWaitingQueue()
	return &manager, nil
}

func (tm *taskManager) RegisterTasks(tasks ...DelayTask) {
	for _, item := range tasks {
		tm.registeredTasks[item.GetName()] = item
	}
}

func (tm *taskManager) Delay(task DelayTask, delay time.Duration, args ...interface{}) {
	tm.broker.AddTask(task, delay, args)
}

func (tm *taskManager) HandleWaitingQueue() {
	tm.broker.HandleWaitingQueue()
}
