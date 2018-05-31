package tasks_test

import (
	"fmt"
	"github.com/dima-kov/go-tasks/tasks"
	"time"
)

type DeleteById struct {
	Name    string
	Manager tasks.TaskManager
}

func (d DeleteById) GetName() string {
	return d.Name
}

func (d DeleteById) Delay(t time.Duration, args ...interface{}) {
	d.Manager.Delay(d, t, args)
}

func (d DeleteById) RunTask(args ...interface{}) {
	//	business logic for deleting
	fmt.Println("Run task. \nDeleting user by id...")
	fmt.Println("Deleted!")
}