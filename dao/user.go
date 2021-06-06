package dao

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID int
	Name string
	CreatedTime time.Time
}


func (u *User)Find(db *gorm.DB,id string) *User {
	user := &User{}
	err := db.Where("id", id).Find(user).Error
	if err != nil{
		return nil
	}
	return user
}