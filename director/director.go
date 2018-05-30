package director

import (
	"github.com/dima-kov/go-tasks/tasks"
	"sync"
	"time"
)

type TasksDirector struct {
	tasks []tasks.PeriodicalTask
}

func NewTasksDirector(t ...tasks.PeriodicalTask) TasksDirector {
	return TasksDirector{tasks: t}
}

func (td *TasksDirector) RegiкsterTask(task tasks.PeriodicalTask) {
	td.tasks = append(td.tasks, task)
}

func (td *TasksDirector) Run() {
	var wg sync.WaitGroup
	wg.Add(len(td.tasks))
	for _, task := range td.tasks {
		go td.executeTask(task)
	}
	wg.Wait()
}

func (td *TasksDirector) executeTask(task tasks.PeriodicalTask) {
	for {
		time.Sleep(task.GetRunEvery())
		task.Execute()
	}
}
