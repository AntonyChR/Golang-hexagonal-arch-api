package middlewares

import (
	"chat/infrastructure/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuhtRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt, err := ctx.Cookie("jwt")
		errorMsg := "Invalid jwt"
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorMsg)
			return
		}

		userId, err := security.ValidateJWT("id", jwt)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorMsg)
			return
		}
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
