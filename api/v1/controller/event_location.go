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
	var req model.EventLocationGetReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
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

	var req model.EventLocationPostReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventLocationPost(c, model.EventLocationPostIn{
		Event:          *req.Event,
		Country:        *req.Country,
		State:          *req.State,
		City:           *req.City,
		Street:         *req.Street,
		Number:         *req.Number,
		AdditionalInfo: *req.AdditionalInfo,
		MapsURL:        *req.MapsURL,
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

	var req model.EventLocationUpdateReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.EventLocationUpdate(c, model.EventLocationUpdateIn{
		ID:                     id,
		EventLocationUpdateReq: req,
	}))
}
