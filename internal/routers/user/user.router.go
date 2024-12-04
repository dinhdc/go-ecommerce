package user

import (
	"github.com/dinhdc/go-ecommerce/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHandlerNonDependencies := controllers.NewUserController(us)
	userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}
	// private router
	userRouterPrivate := Router.Group("/user")

	{
		userRouterPrivate.GET("/info")
	}
}
