package conf

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DbUsername = "root"
const DbPassword = "123456"
const DbName = "more"
const DbHost = "127.0.0.1"
const DbPort = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DbUsername + ":" + DbPassword + "@tcp" + "(" + DbHost + ":" + DbPort + ")/" + DbName + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	}

	return db
}

var Rdb *redis.Client
var Ctx = context.Background()

func RedisInit() *redis.Client {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("connect redis failed")
		return nil
	}
	fmt.Printf("redis ping result: %s\n", pong)
	Rdb = RedisClient
	return Rdb
}
