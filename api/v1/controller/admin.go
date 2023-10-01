package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) AdminLogin() {
	c.Group.POST("/admin/login", adminLogin)
}

func adminLogin(c *gin.Context) {
	req := model.AdminLoginReq{
		Name:     pkg.NewVarchar(20, true),
		Password: pkg.NewVarchar(20, true),
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	in := model.AdminLoginIn{
		Name:     *req.Name.Value,
		Password: *req.Password.Value,
	}

	c.JSON(service.AdminLogin(c, in))
}

func (c Controllers) AdminRegister() {
	c.Group.POST("/admin/register", pkg.RequireAdminSession, adminRegister)
}

func adminRegister(c *gin.Context) {
	adminID := c.GetString("admin-id")
	if adminID == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	req := model.AdminRegisterReq{
		Name:     pkg.NewVarchar(20, true),
		Email:    pkg.NewVarchar(20, true),
		Password: pkg.NewVarchar(20, true),
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	in := model.AdminRegisterIn{
		Name:     *req.Name.Value,
		Email:    *req.Email.Value,
		Password: *req.Password.Value,
		AdminID:  adminID,
	}

	c.JSON(service.AdminRegister(c, in))
}
