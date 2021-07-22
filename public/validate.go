package public

import (
	"Initial/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator2 "github.com/go-playground/validator/v10"
)

func InitValidator() {
	var c *gin.Context
	v, ok := binding.Validator.Engine().(*validator2.Validate)
	if ok {
		// 自定义验证方法
		err := v.RegisterValidation("my_validator", checkName)
		if err != nil {
			middleware.FailResponse(c, 10004, "Field is not start ")
		}
	}
}


func checkName(fl validator2.FieldLevel) bool {
	field := fl.Field().String()
	fmt.Print(field)
	return false
}
