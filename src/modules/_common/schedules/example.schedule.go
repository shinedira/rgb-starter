package schedules

import (
	"fmt"
	"rgb/starter/pkg/scheduler"
	"time"
)

type ExampleSchedule struct{}

func (s *ExampleSchedule) Signature() string {
	return "example:schedule"
}

func (s *ExampleSchedule) Schedule() string {
	return scheduler.EveryMinute()
}

func (s *ExampleSchedule) Handle() {
	fmt.Println("Cron job executed at:", time.Now())
}
