package initialize

import (
	"github.com/dinhdc/go-ecommerce/global"
	"github.com/dinhdc/go-ecommerce/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	adminRouter := routers.RouterGroupApi.Admin
	userRouter := routers.RouterGroupApi.User
	mainGroup := r.Group("/v1/2024")
	{
		mainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitProductRouter(mainGroup)
		userRouter.InitUserRouter(mainGroup)
	}
	{
		adminRouter.InitAdminRouter(mainGroup)
		adminRouter.InitUserRouter(mainGroup)
	}
	return r
}
