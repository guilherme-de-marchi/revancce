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
	var req model.AdminLoginReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	in := model.AdminLoginIn{
		Name:     *req.Name,
		Password: *req.Password,
	}

	c.JSON(service.AdminLogin(c, in))
}

func (c Controllers) AdminRegister() {
	c.Group.POST("/admin/register", pkg.RequireAdminSession, adminRegister)
}

func adminRegister(c *gin.Context) {
	var req model.AdminRegisterReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorMsg(err.Error()))
		return
	}

	id := c.GetString("id")
	if id == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	in := model.AdminRegisterIn{
		Name:     *req.Name,
		Email:    *req.Email,
		Password: *req.Password,
		ID:       id,
	}

	c.JSON(service.AdminRegister(c, in))
}
