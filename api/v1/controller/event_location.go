package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) EventLocationGet() {
	c.Group.GET("/event/location", eventLocationGet)
}

func eventLocationGet(c *gin.Context) {
	req := model.EventLocationGetReq{
		ID:      pkg.NewVarchar(40, false),
		Event:   pkg.NewVarchar(40, false),
		Country: pkg.NewVarchar(20, false),
		State:   pkg.NewVarchar(20, false),
		City:    pkg.NewVarchar(20, false),
		Street:  pkg.NewVarchar(20, false),
		Number:  pkg.NewVarchar(20, false),
		Offset:  pkg.NewInteger(false),
		Page:    pkg.NewInteger(false),
		Limit:   pkg.NewInteger(false),
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

	c.JSON(service.EventLocationGet(c, model.EventLocationGetIn(req)))
}

func (c Controllers) EventLocationPost() {
	c.Group.POST("/event/location", pkg.RequireAdminSession, eventLocationPost)
}

func eventLocationPost(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	req := model.EventLocationPostReq{
		Event:          pkg.NewVarchar(40, true),
		Country:        pkg.NewVarchar(20, true),
		State:          pkg.NewVarchar(20, true),
		City:           pkg.NewVarchar(20, true),
		Street:         pkg.NewVarchar(20, true),
		Number:         pkg.NewVarchar(20, true),
		AdditionalInfo: pkg.NewVarchar(100, true),
		MapsURL:        pkg.NewVarchar(100, true),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventLocationPost(c, model.EventLocationPostIn{
		Event:          *req.Event.Value,
		Country:        *req.Country.Value,
		State:          *req.State.Value,
		City:           *req.City.Value,
		Street:         *req.Street.Value,
		Number:         *req.Number.Value,
		AdditionalInfo: *req.AdditionalInfo.Value,
		MapsURL:        *req.MapsURL.Value,
		AdminID:        adminID,
	}))
}

func (c Controllers) EventLocationDelete() {
	c.Group.DELETE("/event/location/:id", pkg.RequireAdminSession, eventLocationDelete)
}

func eventLocationDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	c.JSON(service.EventLocationDelete(c, model.EventLocationDeleteIn{ID: id}))
}

func (c Controllers) EventLocationUpdate() {
	c.Group.PUT("/event/location/:id", pkg.RequireAdminSession, eventLocationUpdate)
}

func eventLocationUpdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	req := model.EventLocationUpdateReq{
		Event:          pkg.NewVarchar(40, true),
		Country:        pkg.NewVarchar(20, true),
		State:          pkg.NewVarchar(20, true),
		City:           pkg.NewVarchar(20, true),
		Street:         pkg.NewVarchar(20, true),
		Number:         pkg.NewVarchar(20, true),
		AdditionalInfo: pkg.NewVarchar(100, true),
		MapsURL:        pkg.NewVarchar(100, true),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventLocationUpdate(c, model.EventLocationUpdateIn{
		ID:                     id,
		EventLocationUpdateReq: req,
	}))
}
