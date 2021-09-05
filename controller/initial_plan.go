package controller

import (
	"Initial/conf"
	"Initial/dao"
	"Initial/dto"
	"Initial/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PlanController struct{}

func PlanRegister(router *gin.RouterGroup) {
	event := PlanController{}
	router.POST("/create", event.CreatePlan)
}

// @Summary 增加计划
// @Description 增加计划信息
// @Tags 计划信息接口
// @ID /plan/create
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData dto.PlanInput true "params"
// @Success 200 {object} string "success"
// @Router /plan/create  [post]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (eventControl *PlanController) CreatePlan(c *gin.Context) {
	var r dto.PlanInput
	fmt.Print(r)
	err1 := c.ShouldBind(&r)
	fmt.Print(err1)
	if err1 != nil {
		middleware.FailWithDetailed(c, 10001, err1.Error(), "params error")
		return
	}
	plan := dao.Plan{Title: r.Title, User: r.User, Start: r.Start, End: r.End}
	err := dao.CreatePlan(conf.Db, &plan)
	if err != nil {
		middleware.FailResponse(c, 10002, "create failed")
		return
	}
	middleware.SuccessResponseWithData(c, plan)
}
