package model

import (
	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

type CompanyGetReq struct {
	ID     pkg.Varchar `form:"id"`
	Name   pkg.Varchar `form:"name"`
	Offset pkg.Integer `form:"offset"`
	Page   pkg.Integer `form:"page"`
	Limit  pkg.Integer `form:"limit"`
}

func (v CompanyGetReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type CompanyGetIn CompanyGetReq

type CompanyGetOut struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CompanyPostReq struct {
	Name pkg.Varchar `json:"name"`
}

func (v CompanyPostReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type CompanyPostIn struct {
	Name    string
	AdminID string
}

type CompanyDeleteIn struct {
	ID string
}

type CompanyUpdateReq struct {
	Name pkg.Varchar `json:"name"`
}

func (v CompanyUpdateReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type CompanyUpdateIn struct {
	ID string
	CompanyUpdateReq
}
