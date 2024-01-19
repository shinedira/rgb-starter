package scheduler

import (
	sc "rgb/starter/pkg/scheduler"
	"rgb/starter/src/modules/_common/schedules"
	"rgb/starter/src/utils/app"
)

type Scheduler struct{}

func (c *Scheduler) Boot() {
	app.App.Schedule = sc.BootSchedule([]sc.ScheduleEvent{
		&schedules.ExampleSchedule{},
		// other schedule here
	})
}
