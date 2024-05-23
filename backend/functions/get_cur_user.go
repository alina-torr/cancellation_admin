package functions

import (
	"booking/consts"
	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) int64 {
	return c.GetInt64(consts.GIN_USER_ID_KEY)
}
