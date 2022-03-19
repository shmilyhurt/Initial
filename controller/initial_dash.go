package controller

import (
	"Initial/conf"
	"Initial/dao"
	"Initial/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DashController struct{}

func DashRegister(router *gin.RouterGroup) {
	project := DashController{}
	router.GET("/all", project.GetDashAll)
	router.GET("/pro_status", project.GetDashProStatus)
	router.GET("/pro_type", project.GetDashProType)
}

func (dashControl *DashController) GetDashAll(c *gin.Context) {
	var dashAll dao.DashAll
	err := dao.GetDashAll(conf.Db, &dashAll)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	middleware.SuccessResponseWithData(c, dashAll)
}

func (dashControl *DashController) GetDashProStatus(c *gin.Context) {
	var dashProStatus dao.DashProStatus
	err := dao.GetDashProStatus(conf.Db, &dashProStatus)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	middleware.SuccessResponseWithData(c, dashProStatus)
}

func (dashControl *DashController) GetDashProType(c *gin.Context) {
	var dashProType []dao.DashProType
	err := dao.GetDashProType(conf.Db, &dashProType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	result := make([][]int, 0)
	sCount := make([]int, 0)
	mCount := make([]int, 0)
	lCount := make([]int, 0)
	dCount := make([]int, 0)
	for _, v := range dashProType {
		sCount = append(sCount, v.SCount)
		mCount = append(mCount, v.MCount)
		lCount = append(lCount, v.LCount)
		dCount = append(dCount, v.DCount)
	}
	reverse(sCount)
	reverse(mCount)
	reverse(lCount)
	reverse(dCount)
	result = append(result, sCount)
	result = append(result, mCount)
	result = append(result, lCount)
	result = append(result, dCount)
	middleware.SuccessResponseWithData(c, result)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
