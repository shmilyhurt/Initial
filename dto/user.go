package dto

type UserLoginInput struct {
	UserName string`json:"username" form:"username" comment:"姓名" binding:""`
	PassWord string`json:"password" form:"password" comment:"密码"  binding:""`
}