package main

import (
	"github.com/dima-kov/go-tasks/director"
	"github.com/dima-kov/go-tasks/tasks"
	"time"
)

func main() {
	deleteTask := tasks.NewDeleteTask(time.Second * 4)
	dropTask := tasks.NewDropQueryTask(time.Second * 5)
	tasksDirector := director.NewTasksDirector(deleteTask, dropTask)
	tasksDirector.Run()
}
