package model

import (
	"time"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

type AdminLoginReq struct {
	Name     pkg.Varchar `json:"name"`
	Password pkg.Varchar `json:"password"`
}

func (v AdminLoginReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type AdminLoginIn struct {
	Name     string
	Password string
}

type AdminLoginOut struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type AdminRegisterReq struct {
	Name     pkg.Varchar `json:"name"`
	Email    pkg.Varchar `json:"email"`
	Password pkg.Varchar `json:"password"`
}

type AdminRegisterIn struct {
	Name     string
	Email    string
	Password string
	AdminID  string
}

func (v AdminRegisterReq) Validate() error {
	return pkg.ValidateStruct(v)
}
