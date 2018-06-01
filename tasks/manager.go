package tasks

import (
	"fmt"
	"github.com/dima-kov/go-tasks/tasks/brokers"
	"github.com/dima-kov/go-tasks/tasks/task"
	"github.com/pkg/errors"
	"time"
)

const WaitingQueueName = "go_tasks"

type TaskManager interface {
	Delay(task task.Task, delay time.Duration, args ...interface{})
	RegisterTasks(tasks ...task.Task)
	HandleWaitingQueue()
}

type taskManager struct {
	broker          brokers.Broker
	registeredTasks map[string]task.Task
}

func NewTaskManager(brokerType, host string, port uint) (TaskManager, error) {
	var broker brokers.Broker
	switch brokerType {
	case "redis":
		broker = brokers.NewRedisBroker(host, port, "")
	default:
		return nil, errors.New("undefined broker type: " + brokerType)
	}
	manager := taskManager{broker: broker, registeredTasks: map[string]task.Task{}}
	manager.HandleWaitingQueue()
	fmt.Println("after handle start")
	return &manager, nil
}

func (tm *taskManager) RegisterTasks(tasks ...task.Task) {
	for _, item := range tasks {
		tm.registeredTasks[item.GetName()] = item
	}
}

func (tm *taskManager) Delay(task task.Task, delay time.Duration, args ...interface{}) {
	tm.broker.AddTask(task, delay, args)
}

func (tm *taskManager) HandleWaitingQueue() {
	tm.broker.HandleWaitingQueue()
}
