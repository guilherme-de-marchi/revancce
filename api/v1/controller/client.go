package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) ClientGet() {
	c.Group.GET("/client", pkg.RequireAdminSession, clientGet)
}

func clientGet(c *gin.Context) {
	req := model.ClientGetReq{
		ID:     pkg.NewVarchar(40, false),
		Name:   pkg.NewVarchar(20, false),
		Email:  pkg.NewVarchar(30, false),
		CPF:    pkg.NewVarchar(20, false),
		Phone:  pkg.NewVarchar(20, false),
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

	c.JSON(service.ClientGet(c, model.ClientGetIn(req)))
}

func (c Controllers) ClientPost() {
	c.Group.POST("/client", pkg.RequireAdminSession, clientPost)
}

func clientPost(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	req := model.ClientPostReq{
		Name:  pkg.NewVarchar(20, true),
		Email: pkg.NewVarchar(30, true),
		CPF:   pkg.NewVarchar(20, true),
		Phone: pkg.NewVarchar(20, true),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.ClientPost(c, model.ClientPostIn{
		Name:    *req.Name.Value,
		Email:   *req.Email.Value,
		CPF:     *req.CPF.Value,
		Phone:   *req.Phone.Value,
		AdminID: adminID,
	}))
}

func (c Controllers) ClientDelete() {
	c.Group.DELETE("/client/:id", pkg.RequireAdminSession, clientDelete)
}

func clientDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	c.JSON(service.ClientDelete(c, model.ClientDeleteIn{ID: id}))
}

func (c Controllers) ClientUpdate() {
	c.Group.PUT("/client/:id", pkg.RequireAdminSession, clientUpdate)
}

func clientUpdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	req := model.ClientUpdateReq{
		Name:  pkg.NewVarchar(20, false),
		Email: pkg.NewVarchar(30, false),
		CPF:   pkg.NewVarchar(20, false),
		Phone: pkg.NewVarchar(20, false),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.ClientUpdate(c, model.ClientUpdateIn{
		ID:              id,
		ClientUpdateReq: req,
	}))
}
