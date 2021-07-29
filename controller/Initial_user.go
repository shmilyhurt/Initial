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
	router.DELETE("/delete", user.DelUser)
	router.PATCH("/patch", user.PatchUser)
}

// @Summary 增加用户
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
	middleware.SuccessResponseWithData(c, user)
}


// @Summary 删除用户
// @Description 删除单个用户
// @Tags 用户信息接口
// @ID /user/delete
// @Accept  json
// @Produce  json
// @Success 200 {object} string "success"
// @Router /user/delete  [delete]
// @Param id query string true "id"
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (userControl *UserController) DelUser(c *gin.Context) {
	var user dao.User
	id := c.Query("id")
	err := dao.GetUser(conf.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err1 := dao.DeleteUser(conf.Db, &user, id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	middleware.SuccessResponseWithMessage(c, "delete success")
}


// @Summary 修改用户
// @Description 修改用户信息
// @Tags 用户信息接口
// @ID /user/patch
// @Accept  multipart/form-data
// @Produce  json
// @Param id query string true "id"
// @Param name formData dto.UserLoginInput false "name"
// @Success 200 {object} string "success"
// @Router /user/patch  [patch]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (userControl *UserController) PatchUser(c *gin.Context) {
	var r dto.UserLoginInput
	var user dao.User
	id := c.Query("id")
	err2 := dao.GetUser(conf.Db, &user, id)
	if err2 != nil {
		middleware.FailWithDetailed(c, 10001, err2.Error(), "params error")
		return
	}
	err1 := c.ShouldBind(&r)
	if err1 != nil {
		middleware.FailWithDetailed(c, 10001, err1.Error(), "params error")
		return
	}
	if r.UserName != "" {
		user.UserName =  r.UserName
	}
	if r.PassWord != "" {
		user.Password = r.PassWord
	}
	conf.Db.UpdateColumns(&user)
	middleware.SuccessResponseWithData(c, user)
}

