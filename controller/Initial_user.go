package controller

import (
	"Initial/conf"
	"Initial/dao"
	"Initial/dto"
	"Initial/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func UserRegister(router *gin.RouterGroup) {
	user := UserController{}
	router.GET("/users", user.GetUserList)
	router.POST("/create", user.CreateUser)
}

// @Summary 用户信息
// @Description 增加用户信息
// @Tags 用户信息接口
// @ID /user/create
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData dto.UserLoginInput true "name"
// @Success 200 {object} string "success"
// @Router /user/create  [post]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (userControl *UserController) CreateUser(c *gin.Context) {
	var r dto.UserLoginInput
	err1 := c.ShouldBind(&r)
	if err1 != nil {
		middleware.FailWithDetailed(c, 10001, err1.Error(), "params error")
		return
	}
	user := dao.User{UserName: r.UserName, Password: r.PassWord}
	err := dao.CreateUser(conf.Db, &user)
	if err != nil {
		middleware.FailResponse(c, 10002, "create failed")
		return
	}
	middleware.SuccessResponseWithData(c, user)
}

// @Summary 用户信息
// @Description 获取用户信息
// @Tags 用户信息接口
// @ID /user/index
// @Accept  json
// @Produce  json
// @Success 200 {object} string "success"
// @Router /user/users  [get]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (userControl *UserController) GetUserList(c *gin.Context) {
	var user []dao.User
	err := dao.FindAllUser(conf.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
