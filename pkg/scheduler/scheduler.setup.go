package scheduler

import (
	"github.com/gofiber/fiber/v3/log"
	"github.com/robfig/cron/v3"
)

type ScheduleEvent interface {
	Signature() string
	Schedule() string
	Handle()
}

func BootSchedule(schedules []ScheduleEvent) *Schedule {
	cron := cron.New()
	for _, schedule := range schedules {
		_, err := cron.AddFunc(schedule.Schedule(), schedule.Handle)
		if err != nil {
			log.Fatalf("cannot register schedule %s: %v", schedule.Signature(), err)
		}

		log.Infof("schedule %s registered", schedule.Signature())
	}

	return &Schedule{cron}
}
