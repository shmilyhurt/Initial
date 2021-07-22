package dao

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  *int       `json:"is_delete" gorm:"column:is_delete; default:1" description:"删除"`
}

func (u *User) TableName() string {
	return "initial_user"
}


func LoginCheck(db *gorm.DB, u *User) (User, error) {
	var user User
	err := db.Where("user_name = ? AND password = ?", u.UserName, u.Password).Find(&user).Error
	return user, err
}


func (u *User)BeforeSave(db *gorm.DB) (err error){
	u.CreatedAt = time.Now()
	return
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
