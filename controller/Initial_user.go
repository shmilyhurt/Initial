package controller

import (
	"Initial/conf"
	"Initial/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	db *gorm.DB
}

func New() *UserController {
	db := conf.InitDb()
	return &UserController{db: db}
}

func UserRegister(router *gin.RouterGroup){
	user := UserController{}
	router.GET("/index", user.GetUserMessages)
}


// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /user/index
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /user/index  [get]
func (user *UserController)GetUserMessages(c *gin.Context){
	users := New()
	userInfo := (&dao.User{}).Find(users.db,"3")
	c.JSON(http.StatusOK, userInfo)
}