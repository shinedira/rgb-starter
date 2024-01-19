package main

import (
	configApp "rgb/starter/src/config/app"
	"rgb/starter/src/config/database"
	"rgb/starter/src/config/handler"
	"rgb/starter/src/config/scheduler"
	"rgb/starter/src/utils/app"
	"rgb/starter/src/utils/contracts"
)

func main() {
	go func() {
		app.App.Schedule.Run()
		app.App.Handler.Listen(app.App.Config.Env.GetString("APP_HOST"))
	}()

	select {}
}

func init() {
	app.App = &app.Application{}

	configs := []contracts.Config{
		&configApp.Config{},
		&database.Database{},
		&handler.Handler{},
		&scheduler.Scheduler{},
	}

	for _, config := range configs {
		config.Boot()
	}
}
