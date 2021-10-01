package dto

import "time"

type ProjectInput struct {
	Name   string `json:"name" form:"name" comment:"名字" binding:""`
	Type   string `json:"type" form:"type" comment:"类型" binding:""`
	Status string `json:"status" form:"status" comment:"状态" binding:""`
}

type ProjectListInput struct {
	Info     string `json:"info" form:"info" comment:"查找信息" validate:""`
	PageSize int    `json:"page_size" form:"page_size" comment:"页数" validate:"required,min=1,max=999"`
	PageNo   int    `json:"page_no" form:"page_no" comment:"页码" validate:"required,min=1,max=999"`
}

type ProjectListItemOutput struct {
	Id        int       `json:"id" gorm:"primary_key"`
	User      int       `json:"user" form:"user" comment:"用户" binding:""`
	Name      string    `json:"name" form:"name" comment:"名字" binding:""`
	Type      string    `json:"type" form:"type" comment:"类型" binding:""`
	Status    string    `json:"status" form:"status" comment:"状态" binding:""`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

type ProjectListOutput struct {
	List  []ProjectListItemOutput `json:"list" form:"list" `
	Total int64               `json:"total" form:"total"`
}


type ProjectsItem struct{
	Id        int       `json:"id" gorm:"primary_key"`
	User      string    `json:"user" form:"user" comment:"用户" binding:""`
	Name      string    `json:"name" form:"name" comment:"名字" binding:""`
	Type      string    `json:"type" form:"type" comment:"类型" binding:""`
	Status    string    `json:"status" form:"status" comment:"状态" binding:""`
	CreatedAt time.Time `json:"created" gorm:"column:create_at" description:"创建时间"`
}

type ProjectsId struct {
	Ids  string `json:"ids"`
}