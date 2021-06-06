package conf

import (
	"fmt"
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
	dsn := DbUsername +":"+ DbPassword +"@tcp"+ "(" + DbHost + ":" + DbPort +")/" + DbName + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	}

	return db
}
