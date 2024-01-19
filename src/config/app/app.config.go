package app

import (
	"rgb/starter/pkg/config"
	"rgb/starter/src/utils/app"
)

type Config struct{}

func (c *Config) Boot() {
	app.App.Config = config.BootConfig()

	c.RegisterEnv()
}

func (c *Config) RegisterEnv() {
	env := app.App.Config.Env

	env.Set("app", map[string]any{
		"name": env.Env("APP_NAME", "RGB Starter"),
		"env":  env.Env("APP_ENV", "development"),
		"port": env.Env("APP_PORT", "3000"),
		"host": env.Env("APP_HOST"),
	})
}
