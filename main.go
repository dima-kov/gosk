package main

import (
	"github.com/dima-kov/gosk/delay"
	"github.com/dima-kov/gosk/periodic"
	"github.com/dima-kov/gosk/tasks_test"
	periodicTest "github.com/dima-kov/gosk/tasks_test/periodic"
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

	deleteTempData := periodicTest.NewDeleteTask(4 * time.Second)
	dropQuery := periodicTest.NewDropQueryTask(10 * time.Second)
	periodicManager := periodic.NewPeriodicalTasksManager(deleteTempData, dropQuery)
	periodicManager.Run()

	deleteAllTask.Delay(4*time.Second, 22) // <- usage

	wg.Wait()
}
