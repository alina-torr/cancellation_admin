package middleware

import (
	"booking/consts"
	"booking/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(userService services.UserService) gin.HandlerFunc {

	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		tokenString = strings.Split(tokenString, "Bearer ")[1]

		if tokenString == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "does not have token"})
		} else {
			userId, err := userService.ValidateToken(tokenString)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
				context.Abort()
				return
			}
			context.Set(consts.GIN_USER_ID_KEY, userId)

		}
	}
}
