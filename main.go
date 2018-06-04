package main

import (
	"github.com/dima-kov/go-tasks/delay"
	"github.com/dima-kov/go-tasks/tasks_test"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	taskManager, err := delay.NewTaskManager("redis", "localhost", 6379)
	if err != nil {
		panic(err)
	}
	deleteAllTask := tasks_test.DeleteById{"delete_all", taskManager}
	taskManager.RegisterTasks(deleteAllTask)

	deleteAllTask.Delay(4*time.Second, 22) // <- usage

	wg.Wait()
}
