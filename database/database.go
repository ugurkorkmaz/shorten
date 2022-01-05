package database

import (
	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	BAG *redis.Client
	CTX = context.Background()
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	bag := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	BAG = bag
	DB = db
	CTX = context.Background()
}
