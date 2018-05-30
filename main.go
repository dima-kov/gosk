package main

import (
	"github.com/dima-kov/go-tasks/director"
	"github.com/dima-kov/go-tasks/periodic_tasks"
	"time"
)

func main() {
	deleteTask := periodic_tasks.NewDeleteTask(time.Second * 4)
	dropTask := periodic_tasks.NewDropQueryTask(time.Second * 5)
	tasksDirector := director.NewTasksDirector(deleteTask, dropTask)
	tasksDirector.Run()
}
