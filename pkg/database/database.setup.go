package database

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v3/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type DatabaseDriver interface {
	GormDialect() gorm.Dialector
	Dsn() string
}

type Database struct {
	*gorm.DB
	Mongo *mongo.Client
	Cache *redis.Client
}

func BootDatabase() *Database {
	return &Database{}
}

func (d *Database) RegisterDatabase(driver DatabaseDriver, config *gorm.Config) {
	db, _ := gorm.Open(driver.GormDialect(), config)

	d.DB = db
}

func (d *Database) RegisterMongo(url string) {
	m, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Info("Error connecting to MongoDB")
	}

	d.Mongo = m
}

func (d *Database) RegisterCache(opts *redis.Options) {
	cache := redis.NewClient(opts)

	_, err := cache.Ping().Result()
	if err != nil {
		log.Info("Error connecting to Redis")
	}

	d.Cache = cache
}
