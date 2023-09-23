package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/controllers"
)

type group struct {
	Group *gin.RouterGroup
}

func Set(g *gin.RouterGroup) {
	controllers.Set(group{Group: g})
}
