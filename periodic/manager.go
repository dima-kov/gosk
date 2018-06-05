package periodic

import (
	"time"
)

type periodicalTasksManager struct {
	tasks []PeriodicalTask
}

func NewPeriodicalTasksManager(t ...PeriodicalTask) periodicalTasksManager {
	return periodicalTasksManager{tasks: t}
}

func (td *periodicalTasksManager) RegisterTask(task PeriodicalTask) {
	td.tasks = append(td.tasks, task)
}

func (td *periodicalTasksManager) Run() {
	for _, task := range td.tasks {
		go td.executeTask(task)
	}
}

func (td *periodicalTasksManager) executeTask(task PeriodicalTask) {
	for {
		time.Sleep(task.GetRunEvery())
		task.Execute()
	}
}
