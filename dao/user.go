package dao

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          int
	Name        string
	CreatedTime time.Time
}

//get users
func FindAllUser(db *gorm.DB, user *[]User) (err error) {
	err = db.Find(&user).Error
	if err != nil {
		return err
	}
	return nil
}

//create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
