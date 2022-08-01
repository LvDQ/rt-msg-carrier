package router

import (
	"rt-msg-carrier/apis"
	"rt-msg-carrier/configs"
	docs "rt-msg-carrier/docs"
	middleware "rt-msg-carrier/middleware/access_log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter index
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// debug mode or release mode
	cfg := configs.Get()
	gin.SetMode(cfg.GinMode())
	// add access log middleware
	r.Use(middleware.LoggerToFile())
	r.Use(gin.Recovery())

	r.GET("/welcome", apis.Welcome)

	// init gin-swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// user := r.Group("/user")
	// {
	// 	user.GET("/createUser")
	// }

	SetupWsRouter(r)

	return r
}
