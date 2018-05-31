package main

import (
	"github.com/dima-kov/go-tasks/tasks"
	"github.com/dima-kov/go-tasks/tasks_test"
	"time"
)

func main() {
	taskManager, err := tasks.NewTaskManager("redis", "localhost", 6379)
	if err != nil {
		panic(err)
	}
	deleteAllTask := tasks_test.DeleteById{"delete_all", taskManager}
	deleteAllTask.Delay(4*time.Hour, 22)
}
