package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) CompanyGet() {
	c.Group.GET("/company", companyGet)
}

func companyGet(c *gin.Context) {
	req := model.CompanyGetReq{
		ID:     pkg.NewVarchar(40, false),
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

	c.JSON(service.CompanyGet(c, model.CompanyGetIn(req)))
}

func (c Controllers) CompanyPost() {
	c.Group.POST("/company", pkg.RequireAdminSession, companyPost)
}

func companyPost(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	req := model.CompanyPostReq{
		Name: pkg.NewVarchar(20, true),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.CompanyPost(c, model.CompanyPostIn{
		Name:    *req.Name.Value,
		AdminID: adminID,
	}))
}

func (c Controllers) CompanyDelete() {
	c.Group.DELETE("/company/:id", pkg.RequireAdminSession, companyDelete)
}

func companyDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	c.JSON(service.CompanyDelete(c, model.CompanyDeleteIn{ID: id}))
}

func (c Controllers) CompanyUpdate() {
	c.Group.PUT("/company/:id", pkg.RequireAdminSession, companyUpdate)
}

func companyUpdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg("param 'id' is empty"))
		return
	}

	req := model.CompanyUpdateReq{
		Name: pkg.NewVarchar(20, false),
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	c.JSON(service.CompanyUpdate(c, model.CompanyUpdateIn{
		ID:               id,
		CompanyUpdateReq: req,
	}))
}
