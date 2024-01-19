# Golang Starter Minimal

## How to Run
- make env first or you can do it with this comment `cp .env.example .env` adjust env file with your local machine env
- and you just can type this comment to run the app `go run main,go` or `go run .`

## Scheduler
this module will use when you have a repeat job to do on the spesific time, how to use this module

you just need to create file on `src/modules/scheduler`, this is the example code of schedule

```go 
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

```

and register the schedule in `src/config/scheduler/scheduler.config.go` 
```go
func (c *Scheduler) Boot() {
	app.App.Schedule = sc.BootSchedule([]sc.ScheduleEvent{
		&schedules.ExampleSchedule{},
		// other schedule here
	})
}
```

to run scheduler on main.go
```go
app.App.Schedule.Run()
```
