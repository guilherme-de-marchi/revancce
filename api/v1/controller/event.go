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

func (c Controllers) EventPost() {
	c.Group.POST("/event", pkg.RequireAdminSession, eventPost)
}

func eventPost(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	var req model.EventPostReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventPost(c, model.EventPostIn{
		Name:    *req.Name,
		Company: *req.Company,
		AdminID: adminID,
	}))
}

func (c Controllers) EventDelete() {
	c.Group.DELETE("/event/:id", pkg.RequireAdminSession, eventDelete)
}

func eventDelete(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	c.JSON(service.EventDelete(c, model.EventDeleteIn{ID: id, AdminID: adminID}))
}
