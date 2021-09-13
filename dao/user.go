package dao

import (
	"Initial/dto"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete; default:1" description:"删除"`
}

func (u *User) TableName() string {
	return "initial_user"
}

func LoginCheck(db *gorm.DB, u *User) (*gorm.DB, User) {
	var user User
	result := db.Where("user_name = ? AND password = ?", u.UserName, u.Password).Find(&user)
	return result, user
}

//get users
func FindAllUser(db *gorm.DB, user *[]User, count *int64, params *dto.UserListInput) (err error) {
	info := params.Info
	pageNo := params.PageNo
	pageSize := params.PageSize
	offset := (pageNo - 1) * pageSize
	query := db.Limit(pageSize).Offset(offset).Where("is_delete", 1)
	if len(info) != 0 {
		err = query.Where("id like ?", "%"+info+"%").Or("user_name like ?", "%"+info+"%").Find(&user).Count(count).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		err = query.Find(&user).Count(count).Error
		if err != nil {
			return err
		}
		return nil
	}
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

//delete user
func DeleteUser(db *gorm.DB, User *User) (err error) {
	User.IsDelete = 0
	db.Model(User).Update("is_delete", 0)
	return nil
}
