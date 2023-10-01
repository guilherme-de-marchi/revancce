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
	req := model.EventGetReq{
		ID:      pkg.NewVarchar(40, false),
		Name:    pkg.NewVarchar(20, false),
		Company: pkg.NewVarchar(40, false),
		Offset:  pkg.NewInteger(false),
		Page:    pkg.NewInteger(false),
		Limit:   pkg.NewInteger(false),
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventGet(c, model.EventGetIn{
		ID:      req.ID.Value,
		Name:    req.Name.Value,
		Company: req.Company.Value,
		Offset:  req.Offset.Value,
		Page:    req.Page.Value,
		Limit:   req.Limit.Value,
	}))
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

	req := model.EventPostReq{
		Name:    pkg.NewVarchar(20, false),
		Company: pkg.NewVarchar(40, false),
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventPost(c, model.EventPostIn{
		Name:    *req.Name.Value,
		Company: *req.Company.Value,
		AdminID: adminID,
	}))
}

func (c Controllers) EventDelete() {
	c.Group.DELETE("/event/:id", pkg.RequireAdminSession, eventDelete)
}

func eventDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	c.JSON(service.EventDelete(c, model.EventDeleteIn{ID: id}))
}

func (c Controllers) EventUpdate() {
	c.Group.PUT("/event/:id", pkg.RequireAdminSession, eventUpdate)
}

func eventUpdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	req := model.EventUpdateReq{
		Name:    pkg.NewVarchar(20, false),
		Company: pkg.NewVarchar(40, false),
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventUpdate(c, model.EventUpdateIn{
		ID:      id,
		Name:    req.Name.Value,
		Company: req.Company.Value,
	}))
}
