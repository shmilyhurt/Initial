package dao

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	Id int `json:"id" gorm:"primary_key"`
	User int `json:"user" gorm:"column:user;not null"`
	Name string `json:"name" gorm:"column:name;size:256"`
	Type string  `json:"type" gorm:"column:type;size:64"`
	Status string `json:"status" gorm:"column:status;size:64;default:S"`
	Created   time.Time `json:"created" gorm:"autoCreateTime;column:created"`
	IsDelete int `json:"is_delete" gorm:"column:is_delete;default:1"`
}

func (p *Project) TableName() string {
	return "initial_project"
}

//get projects
func FindAllProject(db *gorm.DB, project *[]Project) (err error) {
	err = db.Find(&project).Error
	if err != nil {
		return err
	}
	return nil
}

//create a project
func CreateProject(db *gorm.DB, Project *Project) (err error) {
	err = db.Create(Project).Error
	if err != nil {
		return err
	}
	return nil
}

//get project by id
func GetProject(db *gorm.DB, Project *Project, id string) (err error) {
	err = db.Where("id = ?", id).First(Project).Error
	if err != nil {
		return err
	}
	return nil
}

//delete project
func DeleteProject(db *gorm.DB, Project *Project, id string) (err error) {
	db.Where("id = ?", id).Delete(Project)
	return nil
}
