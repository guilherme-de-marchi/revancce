package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) EventGet() {
	c.Group.GET("/event", eventGet)
}

func eventGet(c *gin.Context) {
	var req model.EventGetReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventGet(c, model.EventGetIn(req)))
}
