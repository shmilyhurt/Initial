package dto

type PlanInput struct {
	User  int    `json:"user" form:"user" comment:"用户id" binding:""`
	Title string `json:"title" form:"title" comment:"标题" binding:""`
	Start string `json:"start" form:"start" comment:"开始时间" binding:"" gorm:"type:date;column:start"  `
	End   string `json:"end" form:"end" comment:"结束时间" binding:"" gorm:"type:date;column:end"`
}

type PlanListItemOutput struct {
	Id    int    `json:"id" gorm:"primary_key"`
	User  int    `json:"user" gorm:"column:user"`
	Title string `json:"title" gorm:"column:title"`
	Start string `json:"start" gorm:"type:date;column:start"`
	End   string `json:"end" gorm:"type:date;column:end"`
}

type PlanListOutput struct {
	List []PlanListItemOutput `json:"list" form:"list"`
}
