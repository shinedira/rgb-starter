package database

import (
	"rgb/starter/pkg/database"
	"rgb/starter/src/config/database/drivers"
	"rgb/starter/src/utils/app"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Database struct{}

func (c *Database) Boot() {
	c.RegisterEnv()
	env := app.App.Config.Env
	db := database.BootDatabase()

	// register driver
	db.RegisterDatabase(&drivers.PostgresDriver{}, &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		// other config here
	})

	// register mongo
	db.RegisterMongo(env.GetString("database.mongo"))

	// register redis
	db.RegisterCache(&redis.Options{
		Addr:     env.GetString("database.redis.host"),
		Password: env.GetString("database.redis.password"),
		DB:       env.GetInt("database.redis.database"),
	})

	app.App.Database = db
}

func (c *Database) RegisterEnv() {
	// set env here
	env := app.App.Config.Env

	env.Set("database", map[string]any{
		"pgsql": map[string]any{
			"host":     env.Env("DB_HOST", "127.0.0.1"),
			"port":     env.Env("DB_PORT", "3306"),
			"database": env.Env("DB_DATABASE", "database"),
			"username": env.Env("DB_USERNAME", ""),
			"password": env.Env("DB_PASSWORD", ""),
		},
		"mongo": env.Env("MONGO_DB_URI", "mongodb://localhost:27017"),
		"redis": map[string]any{
			"host":     env.Env("REDIS_HOST", ""),
			"password": env.Env("REDIS_PASSWORD", ""),
			"port":     env.Env("REDIS_PORT", 6379),
			"database": env.Env("REDIS_DB", 0),
			"prefix":   env.Env("REDIS_PREFIX", "rgb"),
		},
	})
}
