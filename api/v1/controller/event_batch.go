package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) EventBatchGet() {
	c.Group.GET("/event/batch", eventBatchGet)
}

func eventBatchGet(c *gin.Context) {
	req := model.EventBatchGetReq{
		ID:              pkg.NewVarchar(40, false),
		Ticket:          pkg.NewVarchar(40, false),
		Number:          pkg.NewInteger(false),
		FromLimitAmount: pkg.NewInteger(false),
		ToLimitAmount:   pkg.NewInteger(false),
		FromLimitTime:   pkg.NewVarchar(30, false),
		ToLimitTime:     pkg.NewVarchar(30, false),
		Opened:          pkg.NewBoolean(false),
		FromPrice:       pkg.NewInteger(false),
		ToPrice:         pkg.NewInteger(false),
		Offset:          pkg.NewInteger(false),
		Page:            pkg.NewInteger(false),
		Limit:           pkg.NewInteger(false),
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

	c.JSON(service.EventBatchGet(c, model.EventBatchGetIn(req)))
}

func (c Controllers) EventBatchPost() {
	c.Group.POST("/event/batch", pkg.RequireAdminSession, eventBatchPost)
}

func eventBatchPost(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	req := model.EventBatchPostReq{
		Ticket:      pkg.NewVarchar(40, true),
		LimitAmount: pkg.NewInteger(true),
		LimitTime:   pkg.NewVarchar(30, true),
		Price:       pkg.NewInteger(true),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventBatchPost(c, model.EventBatchPostIn{
		Ticket:      *req.Ticket.Value,
		LimitAmount: *req.LimitAmount.Value,
		LimitTime:   *req.LimitTime.Value,
		Price:       *req.Price.Value,
		AdminID:     adminID,
	}))
}

func (c Controllers) EventBatchDelete() {
	c.Group.DELETE("/event/batch/:id", pkg.RequireAdminSession, eventBatchDelete)
}

func eventBatchDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	c.JSON(service.EventBatchDelete(c, model.EventBatchDeleteIn{ID: id}))
}

func (c Controllers) EventBatchUpdate() {
	c.Group.PUT("/event/batch/:id", pkg.RequireAdminSession, eventBatchUpdate)
}

func eventBatchUpdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	req := model.EventBatchUpdateReq{
		Ticket:      pkg.NewVarchar(40, false),
		LimitAmount: pkg.NewInteger(false),
		LimitTime:   pkg.NewVarchar(30, false),
		Opened:      pkg.NewBoolean(false),
		Price:       pkg.NewInteger(false),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventBatchUpdate(c, model.EventBatchUpdateIn{
		ID:                  id,
		EventBatchUpdateReq: req,
	}))
}
