package handlers

import "github.com/gin-gonic/gin"

func handleError(c *gin.Context, code int, err error, sendMessage bool) {
	if !sendMessage {
		c.AbortWithStatus(code)
	} else {
		c.AbortWithStatusJSON(code, gin.H{"message": err.Error()})
	}
}
