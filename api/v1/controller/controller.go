package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

type Controllers struct {
	Group *gin.RouterGroup
}

func Setup(g *gin.RouterGroup) {
	pkg.SetupControllers(Controllers{Group: g})
}
