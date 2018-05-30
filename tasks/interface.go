package tasks

import "time"

type PeriodicalTask interface {
	Execute() error
	GetRunEvery() time.Duration
}