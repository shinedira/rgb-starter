package scheduler

import (
	"github.com/robfig/cron/v3"
)

type Schedule struct {
	cron *cron.Cron
}

func (s *Schedule) Run() {
	s.cron.Start()
}

func (s *Schedule) Stop() {
	s.cron.Stop()
}

func EveryMinute() string {
	return "*/1 * * * *"
}

func Hourly() string {
	return "0 * * * *"
}

func Daily() string {
	return "0 0 * * *"
}

func Weekly() string {
	return "0 0 * * 0"
}

func Monthly() string {
	return "0 0 1 * *"
}

func Yearly() string {
	return "0 0 1 1 *"
}
