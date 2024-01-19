package app

import (
	"rgb/starter/pkg/config"
	"rgb/starter/pkg/database"
	"rgb/starter/pkg/handler"
	"rgb/starter/pkg/scheduler"
)

type Application struct {
	Config *config.Config
	// Cache    *cache.Cache
	Database *database.Database
	Handler  *handler.Handler
	Schedule *scheduler.Schedule
}

var (
	App *Application
)
