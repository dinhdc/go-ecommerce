package controllers

import (
	"net/http"

	"github.com/dinhdc/go-ecommerce/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	name := c.DefaultQuery("name", "aino")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"user":    uc.userService.GetInfoUser(),
	})
}
