package dto

type PlanInput struct {
	User  int    `json:"user" form:"user" comment:"用户id" binding:""`
	Title string `json:"title" form:"title" comment:"标题" binding:""`
	Start string `json:"start" form:"start" comment:"开始时间" binding:"" `
	End   string `json:"end" form:"end" comment:"结束时间" binding:""`
}
