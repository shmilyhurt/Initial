package controller

import (
	"Initial/conf"
	"Initial/dao"
	"Initial/dto"
	"Initial/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"net/http"
)

type PlanController struct{}

func PlanRegister(router *gin.RouterGroup) {
	plan := PlanController{}
	router.GET("/plans", plan.GetPlan)
	router.POST("/create", plan.CreatePlan)
	router.DELETE("/delete", plan.DelPlan)
	router.PATCH("/patch", plan.PatchPlan)
}

// @Summary 增加计划
// @Description 增加计划信息
// @Tags 计划信息接口
// @ID /plan/create
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData dto.PlanInput true "params"
// @Success 200 {object} string  "success"
// @Failure 400 {string} string  "Client data error"
// @Failure 500 {string} string  "Server error"
// @Router /plan/create  [post]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (planControl *PlanController) CreatePlan(c *gin.Context) {
	var r dto.PlanInput
	err1 := c.ShouldBind(&r)
	if err1 != nil {
		middleware.FailWithDetailed(c, 400, err1.Error(), "params error")
		return
	}
	start := carbon.Parse(r.Start, carbon.Shanghai).ToDateTimeString()
	end := carbon.Parse(r.End, carbon.Shanghai).ToDateTimeString()
	plan := dao.Plan{Title: r.Title, User: r.User, Start: start, End: end}
	err := dao.CreatePlan(conf.Db, &plan)
	if err != nil {
		middleware.FailResponse(c, 500, "create failed")
		return
	}
	middleware.SuccessResponseWithData(c, plan)
}

// @Summary 计划信息
// @Description 获取计划信息
// @Tags 计划信息接口
// @ID /plan/plans
// @Accept  json
// @Produce  json
// @Success 200 {object} string "success"
// @Failure 400 {string} string  "Client data error"
// @Failure 500 {string} string  "Server error"
// @Param user query string true "user"
// @Router /plan/plans  [get]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (planControl *PlanController) GetPlan(c *gin.Context) {
	var plan []dao.Plan
	user := c.Query("user")
	err := dao.FindUserPlan(conf.Db, &plan, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var outputList []dto.PlanListItemOutput
	for _, item := range plan{
		outputList = append(outputList, dto.PlanListItemOutput{
			Id: item.Id,
			User: item.User,
			Title: item.Title,
			Start: item.Start,
			End: item.End,
		})
	}
	output := dto.PlanListOutput{
		List:  outputList,
	}
	middleware.SuccessResponseWithData(c, output)
}

// @Summary 删除计划
// @Description 删除单个计划
// @Tags 计划信息接口
// @ID /plan/delete
// @Accept  json
// @Produce  json
// @Success 200 {object} string "success"
// @Failure 400 {string} string  "Client data error"
// @Failure 500 {string} string  "Server error"
// @Router /plan/delete  [delete]
// @Param id query string true "id"
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (planControl *PlanController) DelPlan(c *gin.Context) {
	var plan dao.Plan
	id := c.Query("id")
	err := dao.GetPlan(conf.Db, &plan, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err1 := dao.DeletePlan(conf.Db, &plan)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	middleware.SuccessResponseWithMessage(c, "delete success")
}

// @Summary 修改计划
// @Description 修改计划信息
// @Tags 计划信息接口
// @ID /plan/patch
// @Accept  multipart/form-data
// @Produce  json
// @Param id query string true "id"
// @Param name formData dto.PlanInput false "name"
// @Failure 400 {string} string  "Client data error"
// @Failure 500 {string} string  "Server error"
// @Success 200 {object} string "success"
// @Router /plan/patch  [patch]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (planControl *PlanController) PatchPlan(c *gin.Context) {
	var r dto.PlanInput
	var plan dao.Plan
	id := c.Query("id")
	err2 := dao.GetPlan(conf.Db, &plan, id)
	if err2 != nil {
		middleware.FailWithDetailed(c, 500, err2.Error(), "params error")
		return
	}
	err1 := c.ShouldBind(&r)
	if err1 != nil {
		middleware.FailWithDetailed(c, 500, err1.Error(), "params error")
		return
	}
	if r.Title != "" {
		plan.Title = r.Title
	}
	if r.Start != "" {
		plan.Start = carbon.Parse(r.Start, carbon.Shanghai).ToDateTimeString()
	}
	if r.End != "" {
		plan.End = carbon.Parse(r.End, carbon.Shanghai).ToDateTimeString()
	}
	conf.Db.Updates(&plan)
	middleware.SuccessResponseWithData(c, plan)
}
