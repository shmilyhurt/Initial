package controller

import (
	"Initial/conf"
	"Initial/dao"
	"Initial/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

type DashProStatusController struct{}

func DashProStatusRegister(router *gin.RouterGroup) {
	project := DashProStatusController{}
	router.GET("/pro_status", project.GetDashProStatus)
	router.GET("/pro_type", project.GetDashProType)
}

func (dashProStatusControl *DashProStatusController) GetDashProStatus(c *gin.Context) {
	var dashProStatus dao.DashProStatus
	err := dao.GetDashProStatus(conf.Db, &dashProStatus)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	middleware.SuccessResponseWithData(c, dashProStatus)
}

func (dashProStatusControl *DashProStatusController) GetDashProType(c *gin.Context) {
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
	fmt.Println(result)
	middleware.SuccessResponseWithData(c, result)
}

func reverse(a []int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}
