package dao

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	Id int `json:"id" gorm:"primary_key"`
	Pro int `json:"pro"gorm:"column:pro;not null"`
	User int `json:"user"gorm:"column:user;not null"`
	Name string `json:"name" gorm:"column:name;size:256"`
	Type string  `json:"type" gorm:"column:type;size:64"`
	Status string `json:"status"gorm:"column:status;size:64;default:S"`
	Created   time.Time `json:"created"gorm:"autoCreateTime;column:created"`
	IsDelete int `json:"is_delete" gorm:"column:is_delete;default:1"`
}

func (e *Event) TableName() string {
	return "initial_event"
}

//get projects
func FindAllEvent(db *gorm.DB, event *[]Event) (err error) {
	err = db.Find(&event).Error
	if err != nil {
		return err
	}
	return nil
}

//create a project
func CreateEvent(db *gorm.DB, Event *Event) (err error) {
	err = db.Create(Event).Error
	if err != nil {
		return err
	}
	return nil
}

//get project by id
func GetEvent(db *gorm.DB, Event *Event, id string) (err error) {
	err = db.Where("id = ?", id).First(Event).Error
	if err != nil {
		return err
	}
	return nil
}

//delete project
func DeleteEvent(db *gorm.DB, Event *Event, id string) (err error) {
	db.Where("id = ?", id).Delete(Event)
	return nil
}
