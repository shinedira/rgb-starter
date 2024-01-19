package drivers

import (
	"fmt"
	"rgb/starter/src/utils/app"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDriver struct{}

func (p *PostgresDriver) GormDialect() gorm.Dialector {
	return postgres.Open(p.Dsn())
}

func (p *PostgresDriver) Dsn() string {
	config := app.App.Config.Env
	pgsql := config.Get("database.pgsql").(map[string]any)

	// "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		pgsql["host"],
		pgsql["username"],
		pgsql["password"],
		pgsql["database"],
		pgsql["port"],
	)
}
