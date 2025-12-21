package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"chunli/utils/response"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		response.InternalError(c, "Internal server error")
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}

