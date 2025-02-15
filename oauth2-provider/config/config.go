package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/go-redis/redis/v8"
)

var DB *gorm.DB
var RedisClient *redis.Client

func InitDB() {
    var err error
    DB, err = gorm.Open("postgres", "host=localhost user=youruser dbname=oauth2provider sslmode=disable password=yourpassword")
    if err != nil {
        panic("failed to connect database")
    }
}

func InitRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
}