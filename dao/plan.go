package dao

import (
	"gorm.io/gorm"
	"time"
)

type Plan struct {
	Id int `gorm:"primary_key"`
	User int `gorm:"column:user;not null"`
	Title string `gorm:"column:title;not null;size:128"`
	Start string `gorm:"type:date;column:start"`
	End string  `gorm:"type:date;column:end"`
	Created   time.Time `gorm:"autoCreateTime;column:created;type:datetime"`
	IsDelete int `gorm:"column:is_delete;default:1"`
}

func (u *Plan) TableName() string {
	return "initial_plan"
}

//get plans
func FindAllPlan(db *gorm.DB, plan *[]Plan) (err error) {
	err = db.Find(&plan).Error
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
func DeletePlan(db *gorm.DB, plan *Plan, id string) (err error) {
	db.Where("id = ?", id).Delete(plan)
	return nil
}

