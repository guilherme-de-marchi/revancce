package controller

import (
	"net/http"
	"testing"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

func TestAdminLogin_Success(t *testing.T) {
	status, body, err := pkg.NewHTTPRequest(
		http.MethodPost,
		"/api/v1/admin/login",
		model.AdminLoginReq{
			Name:     pkg.Pointer("test2"),
			Password: pkg.Pointer("test2"),
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	if status != http.StatusOK {
		t.Errorf("wrong status code: %v", status)
		t.Log(body)
		return
	}
}

func TestAdminRegister_Success(t *testing.T) {
	status, body, err := pkg.NewHTTPRequest(
		http.MethodPost,
		"/api/v1/admin/register",
		model.AdminRegisterReq{
			Name:     pkg.Pointer("test2"),
			Email:    pkg.Pointer("test2@gmail.com"),
			Password: pkg.Pointer("test2"),
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	if status != http.StatusCreated {
		t.Errorf("wrong status code: %v", status)
		t.Log(body)
		return
	}
}
