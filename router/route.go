package router

import (
	"Initial/controller"
	"Initial/docs"
	"Initial/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// set swagger info
	docs.SwaggerInfo.Title = "Initial swagger api"
	docs.SwaggerInfo.Description = "It is just a server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8880"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.Use(middleware.Cors())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	loginRouter := router.Group("")
	loginRouter.Use(
	)
	{
		controller.LoginRegister(loginRouter)
	}

	userRouter := router.Group("/user")
	userRouter.Use(
		middleware.JwtAuth(),
	)
	{
		controller.UserRegister(userRouter)
	}

	projectRouter := router.Group("/project")
	projectRouter.Use(
		middleware.JwtAuth(),
	)
	{
		controller.ProjectRegister(projectRouter)
	}

	eventRouter := router.Group("/event")
	eventRouter.Use(
		middleware.JwtAuth(),
	)
	{
		controller.EventRegister(eventRouter)
	}


	planRouter := router.Group("/plan")
	planRouter.Use(
		middleware.JwtAuth(),
	)
	{
		controller.PlanRegister(planRouter)
	}

	return router
}
