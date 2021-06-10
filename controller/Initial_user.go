package controller

import (
	"Initial/conf"
	"Initial/dao"
	"Initial/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func UserRegister(router *gin.RouterGroup) {
	user := UserController{}
	router.GET("/index", user.GetUserMessages)
	router.POST("/create", user.CreateUser)
}

// @Summary 用户信息
// @Description 增加用户信息
// @Tags 用户信息接口
// @ID /user/create
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData string true "name"
// @Success 200 {object} string "success"
// @Router /user/create  [post]
func (userControl *UserController) CreateUser(c *gin.Context) {
	var user dao.User
	name := c.PostForm("name")
	if name != "" {
		user.Name = name
	}
	errBindData := c.ShouldBind(&user)
	if errBindData != nil {
		return
	}
	err := dao.CreateUser(conf.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	//c.JSON(http.StatusOK, user)
	middleware.SuccessResponseWithData(c, user)
}

// @Summary 用户信息
// @Description 获取用户信息
// @Tags 用户信息接口
// @ID /user/index
// @Accept  json
// @Produce  json
// @Success 200 {object} string "success"
// @Router /user/index  [get]
func (userControl *UserController) GetUserMessages(c *gin.Context) {
	var user []dao.User
	err := dao.FindAllUser(conf.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
