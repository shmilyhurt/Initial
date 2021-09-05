package controller

import (
	"Initial/conf"
	"Initial/dao"
	"Initial/dto"
	"Initial/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProjectController struct{}

func ProjectRegister(router *gin.RouterGroup) {
	project := ProjectController{}
	router.GET("/projects", project.GetProjectList)
	router.POST("/create", project.CreateProject)
	router.DELETE("/delete", project.DelProject)
	router.PATCH("/patch", project.PatchProject)
}

// @Summary 增加项目
// @Description 增加项目信息
// @Tags 项目信息接口
// @ID /project/create
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData dto.ProjectInput true "name"
// @Success 200 {object} string "success"
// @Router /project/create  [post]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (projectControl *ProjectController) CreateProject(c *gin.Context) {
	var r dto.ProjectInput
	err1 := c.ShouldBind(&r)
	if err1 != nil {
		middleware.FailWithDetailed(c, 10001, err1.Error(), "params error")
		return
	}
	project := dao.Project{Name: r.Name, Type: r.Type, Status: r.Status}
	fmt.Print(project)
	err := dao.CreateProject(conf.Db, &project)
	if err != nil {
		middleware.FailResponse(c, 10002, "create failed")
		return
	}
	middleware.SuccessResponseWithData(c, project)
}

// @Summary 项目信息
// @Description 获取项目信息
// @Tags 项目信息接口
// @ID /project/projects
// @Accept  json
// @Produce  json
// @Success 200 {object} string "success"
// @Router /project/projects  [get]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (projectControl *ProjectController) GetProjectList(c *gin.Context) {
	var project []dao.Project
	var count int64
	params := &dto.ProjectListInput{}
	params.Info = c.Query("Info")
	params.PageSize, _ = strconv.Atoi(c.Query("PageSize"))
	params.PageNo, _ = strconv.Atoi(c.Query("PageNo"))
	err := dao.FindAllProject(conf.Db, &project, &count, params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var outputList []dto.ProjectListItemOutput
	for _, item := range project{
		outputList = append(outputList, dto.ProjectListItemOutput{
			Id: item.Id,
			User: item.User,
			Name: item.Name,
			Type: item.Type,
			Status: item.Status,
			CreatedAt: item.Created,
		})
	}
	output := dto.ProjectListOutput{
		List:  outputList,
		Total: count,
	}
	middleware.SuccessResponseWithData(c, output)
}


// @Summary 删除项目
// @Description 删除单个项目
// @Tags 项目信息接口
// @ID /project/delete
// @Accept  json
// @Produce  json
// @Success 200 {object} string "success"
// @Router /project/delete  [delete]
// @Param id query string true "id"
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (projectControl *ProjectController) DelProject(c *gin.Context) {
	var project dao.Project
	id := c.Query("id")
	err := dao.GetProject(conf.Db, &project, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err1 := dao.DeleteProject(conf.Db, &project)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	middleware.SuccessResponseWithMessage(c, "delete success")
}


// @Summary 修改项目
// @Description 修改项目信息
// @Tags 项目信息接口
// @ID /project/patch
// @Accept  multipart/form-data
// @Produce  json
// @Param id query string true "id"
// @Param name formData dto.ProjectInput false "name"
// @Success 200 {object} string "success"
// @Router /project/patch  [patch]
// @Param token header string true "Insert your access token" default(Bearer <Add access token here>)
func (projectControl *ProjectController) PatchProject(c *gin.Context) {
	var r dto.ProjectInput
	var project dao.Project
	id := c.Query("id")
	err2 := dao.GetProject(conf.Db, &project, id)
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
		project.Name =  r.Name
	}
	if r.Type != "" {
		project.Type = r.Type
	}
	if r.Status != "" {
		project.Status = r.Status
	}
	conf.Db.UpdateColumns(&project)
	middleware.SuccessResponseWithData(c, project)
}

