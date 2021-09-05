package dao

import (
	"Initial/dto"
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
func FindAllProject(db *gorm.DB,project *[]Project, count *int64,  params *dto.ProjectListInput) (err error) {
	info := params.Info
	pageNo := params.PageNo
	pageSize := params.PageSize
	offset := (pageNo - 1) * pageSize
	query := db.Limit(pageSize).Offset(offset).Where("is_delete", 1)
	if len(info) != 0{
		err = query.Where("id like ?", "%" +info + "%").Or("name like ?", "%" +info + "%").Find(&project).Count(count).Error
		if err != nil {
			return err
		}
		return nil
	}else{
		err = query.Find(&project).Count(count).Error
		if err != nil {
			return err
		}
		return nil
	}
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
func DeleteProject(db *gorm.DB, Project *Project) (err error) {
	Project.IsDelete = 0
	db.Model(Project).Update("is_delete", 0)
	return nil
}
