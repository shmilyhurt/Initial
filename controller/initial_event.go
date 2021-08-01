package controller

import (
	"Initial/conf"
	"Initial/dao"
	"Initial/dto"
	"Initial/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventController struct{}

func EventRegister(router *gin.RouterGroup) {
	event := EventController{}
	router.GET("/events", event.GetEventList)
	router.POST("/create", event.CreateEvent)
	router.DELETE("/delete", event.DelEvent)
	router.PATCH("/patch", event.PatchEvent)
}

// @Summary 增加事件
// @Description 增加事件信息
// @Tags 事件信息接口
// @ID /event/create
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData dto.EventInput true "params"
// @Success 200 {object} string "success"
// @Router /event/create  [post]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (eventControl *EventController) CreateEvent(c *gin.Context) {
	var r dto.EventInput
	err1 := c.ShouldBind(&r)
	if err1 != nil {
		middleware.FailWithDetailed(c, 10001, err1.Error(), "params error")
		return
	}
	event := dao.Event{Name: r.Name, Type: r.Type, Status: r.Status}
	err := dao.CreateEvent(conf.Db, &event)
	if err != nil {
		middleware.FailResponse(c, 10002, "create failed")
		return
	}
	middleware.SuccessResponseWithData(c, event)
}

// @Summary 事件信息
// @Description 获取事件信息
// @Tags 事件信息接口
// @ID /event/events
// @Accept  json
// @Produce  json
// @Success 200 {object} string "success"
// @Router /event/events  [get]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (eventControl *EventController) GetEventList(c *gin.Context) {
	var event []dao.Event
	err := dao.FindAllEvent(conf.Db, &event)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	middleware.SuccessResponseWithData(c, event)
}


// @Summary 删除事件
// @Description 删除单个事件
// @Tags 事件信息接口
// @ID /event/delete
// @Accept  json
// @Produce  json
// @Success 200 {object} string "success"
// @Router /event/delete  [delete]
// @Param id query string true "id"
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (eventControl *EventController) DelEvent(c *gin.Context) {
	var event dao.Event
	id := c.Query("id")
	err := dao.GetEvent(conf.Db, &event, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err1 := dao.DeleteEvent(conf.Db, &event, id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	middleware.SuccessResponseWithMessage(c, "delete success")
}


// @Summary 修改事件
// @Description 修改事件信息
// @Tags 事件信息接口
// @ID /event/patch
// @Accept  multipart/form-data
// @Produce  json
// @Param id query string true "id"
// @Param name formData dto.EventInput false "params"
// @Success 200 {object} string "success"
// @Router /event/patch  [patch]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (eventControl *EventController) PatchEvent(c *gin.Context) {
	var r dto.EventInput
	var event dao.Event
	id := c.Query("id")
	err2 := dao.GetEvent(conf.Db, &event, id)
	if err2 != nil {
		middleware.FailWithDetailed(c, 10001, err2.Error(), "params error")
		return
	}
	err1 := c.ShouldBind(&r)
	if err1 != nil {
		middleware.FailWithDetailed(c, 10001, err1.Error(), "params error")
		return
	}
	if r.Name != "" {
		event.Name =  r.Name
	}
	if r.Type != "" {
		event.Type = r.Type
	}
	if r.Status != "" {
		event.Status = r.Status
	}
	conf.Db.UpdateColumns(&event)
	middleware.SuccessResponseWithData(c, event)
}

