package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) EventScheduleGet() {
	c.Group.GET("/event/schedule", eventScheduleGet)
}

func eventScheduleGet(c *gin.Context) {
	req := model.EventScheduleGetReq{
		ID:     pkg.NewVarchar(40, false),
		Event:  pkg.NewVarchar(40, false),
		From:   pkg.NewVarchar(20, false),
		To:     pkg.NewVarchar(20, false),
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
		c.JSON(http.StatusInternalServerError, pkg.ErrorMsg("unable to bind query"))
		return
	}

	c.JSON(service.EventScheduleGet(c, model.EventScheduleGetIn(req)))
}

func (c Controllers) EventSchedulePost() {
	c.Group.POST("/event/schedule", pkg.RequireAdminSession, eventSchedulePost)
}

func eventSchedulePost(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	req := model.EventSchedulePostReq{
		Event:    pkg.NewVarchar(40, true),
		StartsAt: pkg.NewVarchar(20, true),
		EndsAt:   pkg.NewVarchar(20, true),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventSchedulePost(c, model.EventSchedulePostIn{
		Event:    *req.Event.Value,
		StartsAt: *req.StartsAt.Value,
		EndsAt:   *req.EndsAt.Value,
		AdminID:  adminID,
	}))
}

func (c Controllers) EventScheduleDelete() {
	c.Group.DELETE("/event/schedule/:id", pkg.RequireAdminSession, eventScheduleDelete)
}

func eventScheduleDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	c.JSON(service.EventScheduleDelete(c, model.EventScheduleDeleteIn{ID: id}))
}

func (c Controllers) EventScheduleUpdate() {
	c.Group.PUT("/event/schedule/:id", pkg.RequireAdminSession, eventScheduleUpdate)
}

func eventScheduleUpdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	req := model.EventScheduleUpdateReq{
		Event:    pkg.NewVarchar(40, true),
		StartsAt: pkg.NewVarchar(20, false),
		EndsAt:   pkg.NewVarchar(20, false),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventScheduleUpdate(c, model.EventScheduleUpdateIn{
		ID:                     id,
		EventScheduleUpdateReq: req,
	}))
}
