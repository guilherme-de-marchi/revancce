package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAdminSession(c *gin.Context) {
	id, err := GetAdminSession(c, c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.Set("id", id)
	c.Next()
}
