package dto

import "time"

type UserLoginInput struct {
	UserName string`json:"username" form:"username" comment:"姓名" binding:""`
	PassWord string`json:"password" form:"password" comment:"密码"  binding:""`
}

type UserListInput struct {
	Info     string `json:"info" form:"info" comment:"查找信息" validate:""`
	PageSize int    `json:"page_size" form:"page_size" comment:"页数" validate:"required,min=1,max=999"`
	PageNo   int    `json:"page_no" form:"page_no" comment:"页码" validate:"required,min=1,max=999"`
}

type UserListItemOutput struct {
	Id        int     `json:"id" gorm:"primary_key"`
	UserName      string    `json:"user_name" gorm:"column:user_name" description:"名称	"`
	Password      string    `json:"password" gorm:"column:password" description:"密码"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}


type UserListOutput struct {
	List  []UserListItemOutput `json:"list" form:"list" comment:"租户列表"`
	Total int64               `json:"total" form:"total" comment:"租户总数"`
}