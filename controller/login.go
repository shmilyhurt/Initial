package controller

import (
	"Initial/conf"
	"Initial/dao"
	"Initial/dto"
	"Initial/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func LoginRegister(router *gin.RouterGroup) {
	router.POST("/login", Login)
	router.POST("/register", RegisterUser)
}


// @Summary 注册
// @Description 增加用户信息
// @Tags 用户信息接口
// @ID /user/register
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData dto.UserLoginInput true "name"
// @Success 200 {object} string "success"
// @Router /register  [post]
func RegisterUser(c *gin.Context) {
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


// @Summary 登录
// @Description 登录接口
// @Tags 用户信息接口
// @ID /user/Login
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData dto.UserLoginInput true "name"
// @Success 200 {object} string "success"
// @Router /login  [post]
func Login(c *gin.Context) {
	var r dto.UserLoginInput
	err1 := c.ShouldBind(&r)
	if err1 != nil {
		middleware.FailWithDetailed(c, 10001, err1.Error(), "params error")
		return
	}
	user := dao.User{UserName: r.UserName, Password: r.PassWord}
	userInfo, err := dao.LoginCheck(conf.Db, &user)
	if err != nil {
		middleware.FailResponse(c, 10040, "user not exits")
		return
	} else {
		generateToken(c, userInfo)
	}

}

func generateToken(c *gin.Context, user dao.User) {
	j := &middleware.JWT{
		SigningKey: []byte(middleware.JwtConfig{}.SigningKey),
	}

	claims := middleware.CustomClaims{
		ID:       user.Id,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 360000), // 过期时间 一小时
			Issuer:    "lyj",                           //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		middleware.FailResponse(c, 10020, err.Error())
		return
	}

	data := middleware.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}
	middleware.SuccessResponseDetailed(c, data, "login success")
}
