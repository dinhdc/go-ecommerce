package admin

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// userRouterPublic := Router.Group("/admin/user")
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.POST("/otp")
	// }
	// private router
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.GET("/active-user")
	}
}
