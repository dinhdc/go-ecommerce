package controllers

import (
	"github.com/dinhdc/go-ecommerce/internal/service"
	"github.com/dinhdc/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

// type UserController struct {
// 	userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	name := c.DefaultQuery("name", "aino")
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong" + name,
// 		"user":    uc.userService.GetInfoUser(),
// 	})
// }

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}
