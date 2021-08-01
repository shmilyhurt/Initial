package dto

type EventInput struct {
	Name string`json:"name" form:"name" comment:"名字" binding:""`
	Type string`json:"type" form:"type" comment:"类型" binding:""`
	Status string`json:"status" form:"status" comment:"状态" binding:""`
}
