package periodic_tasks

import (
	"fmt"
	"time"
)

type dropQuery struct {
	runEvery time.Duration
}

func NewDropQueryTask(every time.Duration) dropQuery {
	return dropQuery{every}
}

func (d dropQuery) Execute() error {
	//	make deleting here
	fmt.Println("Dropping")
	return nil
}

func (d dropQuery) GetRunEvery() time.Duration {
	return d.runEvery
}
