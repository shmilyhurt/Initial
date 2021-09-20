package dao

import (
	"gorm.io/gorm"
	"time"
)

type Plan struct {
	Id       int       `gorm:"primary_key"`
	User     int       `gorm:"column:user;not null"`
	Title    string    `gorm:"column:title;not null;size:128"`
	Start    string    `gorm:"type:datetime;column:start"`
	End      string    `gorm:"type:datetime;column:end"`
	Created  time.Time `gorm:"autoCreateTime;column:created;type:datetime"`
	IsDelete int       `gorm:"column:is_delete;default:1"`
}

func (u *Plan) TableName() string {
	return "initial_plan"
}

//get plans
func FindUserPlan(db *gorm.DB, plan *[]Plan, user string) (err error) {
	err = db.Where("user = ? AND is_delete = ?", user, 1).Find(&plan).Error
	if err != nil {
		return err
	}
	return nil
}

//create a plan
func CreatePlan(db *gorm.DB, plan *Plan) (err error) {
	err = db.Create(plan).Error
	if err != nil {
		return err
	}
	return nil
}

//delete plan
func DeletePlan(db *gorm.DB, plan *Plan) (err error) {
	db.Model(plan).Update("is_delete", 0)
	return nil
}

//get plan by id
func GetPlan(db *gorm.DB, plan *Plan, id string) (err error) {
	err = db.Where("id = ?", id).First(plan).Error
	if err != nil {
		return err
	}
	return nil
}
