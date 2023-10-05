package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) EventTicketGet() {
	c.Group.GET("/event/ticket", eventTicketGet)
}

func eventTicketGet(c *gin.Context) {
	req := model.EventTicketGetReq{
		ID:     pkg.NewVarchar(40, false),
		Event:  pkg.NewVarchar(40, false),
		Name:   pkg.NewVarchar(20, false),
		Offset: pkg.NewInteger(false),
		Page:   pkg.NewInteger(false),
		Limit:  pkg.NewInteger(false),
	}

	m := make(map[string]string)
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := pkg.BindQuery(m, &req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventTicketGet(c, model.EventTicketGetIn(req)))
}

func (c Controllers) EventTicketPost() {
	c.Group.POST("/event/ticket", pkg.RequireAdminSession, eventTicketPost)
}

func eventTicketPost(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	req := model.EventTicketPostReq{
		Event: pkg.NewVarchar(40, true),
		Name:  pkg.NewVarchar(20, true),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventTicketPost(c, model.EventTicketPostIn{
		Event:   *req.Event.Value,
		Name:    *req.Name.Value,
		AdminID: adminID,
	}))
}

func (c Controllers) EventTicketDelete() {
	c.Group.DELETE("/event/ticket/:id", pkg.RequireAdminSession, eventTicketDelete)
}

func eventTicketDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	c.JSON(service.EventTicketDelete(c, model.EventTicketDeleteIn{ID: id}))
}

func (c Controllers) EventTicketUpdate() {
	c.Group.PUT("/event/ticket/:id", pkg.RequireAdminSession, eventTicketUpdate)
}

func eventTicketUpdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	req := model.EventTicketUpdateReq{
		Event: pkg.NewVarchar(40, false),
		Name:  pkg.NewVarchar(20, false),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventTicketUpdate(c, model.EventTicketUpdateIn{
		ID:                   id,
		EventTicketUpdateReq: req,
	}))
}
