package middlewares

import (
	"github.com/dinhdc/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrInvalidToken, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
