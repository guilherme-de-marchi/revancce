package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/v1/controller"
)

func Setup(e *gin.Engine) {
	controller.Setup(e.Group("/api/v1"))
}
