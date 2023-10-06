package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) ClientTicketGet() {
	c.Group.GET("/client/ticket", pkg.RequireAdminSession, clientTicketGet)
}

func clientTicketGet(c *gin.Context) {
	req := model.ClientTicketGetReq{
		ID:          pkg.NewVarchar(40, false),
		Client:      pkg.NewVarchar(40, false),
		Batch:       pkg.NewVarchar(40, false),
		Transaction: pkg.NewVarchar(40, false),
		From:        pkg.NewVarchar(30, false),
		To:          pkg.NewVarchar(30, false),
		Used:        pkg.NewBoolean(false),
		Offset:      pkg.NewInteger(false),
		Page:        pkg.NewInteger(false),
		Limit:       pkg.NewInteger(false),
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

	c.JSON(service.ClientTicketGet(c, model.ClientTicketGetIn(req)))
}

func (c Controllers) ClientTicketPost() {
	c.Group.POST("/client/ticket", pkg.RequireAdminSession, clientTicketPost)
}

func clientTicketPost(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	req := model.ClientTicketPostReq{
		Client:      pkg.NewVarchar(40, true),
		Batch:       pkg.NewVarchar(40, true),
		Transaction: pkg.NewVarchar(40, true),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.ClientTicketPost(c, model.ClientTicketPostIn{
		Client:      *req.Client.Value,
		Batch:       *req.Batch.Value,
		Transaction: *req.Transaction.Value,
		AdminID:     adminID,
	}))
}

func (c Controllers) ClientTicketDelete() {
	c.Group.DELETE("/client/ticket/:id", pkg.RequireAdminSession, clientTicketDelete)
}

func clientTicketDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	c.JSON(service.ClientTicketDelete(c, model.ClientTicketDeleteIn{ID: id}))
}

func (c Controllers) ClientTicketUpdate() {
	c.Group.PUT("/client/ticket/:id", pkg.RequireAdminSession, clientTicketUpdate)
}

func clientTicketUpdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	req := model.ClientTicketUpdateReq{
		Client:      pkg.NewVarchar(40, false),
		Batch:       pkg.NewVarchar(40, false),
		Transaction: pkg.NewVarchar(40, false),
		Used:        pkg.NewBoolean(false),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.ClientTicketUpdate(c, model.ClientTicketUpdateIn{
		ID:                    id,
		ClientTicketUpdateReq: req,
	}))
}

func (c Controllers) ClientTicketCheckin() {
	c.Group.POST("/client/ticket/check-in", pkg.RequireAdminSession, clientTicketCheckin)
}

func clientTicketCheckin(c *gin.Context) {
	req := model.ClientTicketCheckinReq{
		ID: pkg.NewVarchar(40, true),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.ClientTicketCheckin(c, model.ClientTicketCheckinIn{
		ID: *req.ID.Value,
	}))
}
