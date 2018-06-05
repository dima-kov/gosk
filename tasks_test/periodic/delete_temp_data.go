package periodic

import (
	"fmt"
	"time"
)

type deleteTempData struct {
	runEvery time.Duration
}

func NewDeleteTask(every time.Duration) deleteTempData {
	return deleteTempData{every}
}

func (d deleteTempData) Execute() error {
	//	make deleting here
	fmt.Println("Deleting")
	return nil
}

func (d deleteTempData) GetRunEvery() time.Duration {
	return d.runEvery
}
