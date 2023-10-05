package model

import "github.com/guilherme-de-marchi/revancce/api/pkg"

type EventGetReq struct {
	ID      pkg.Varchar `form:"id"`
	Name    pkg.Varchar `form:"name"`
	Company pkg.Varchar `form:"company"`
	Offset  pkg.Integer `form:"offset"`
	Page    pkg.Integer `form:"page"`
	Limit   pkg.Integer `form:"limit"`
}

func (v EventGetReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventGetIn EventGetReq

type EventGetOut struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
}

type EventPostReq struct {
	Name    pkg.Varchar `json:"name"`
	Company pkg.Varchar `json:"company"`
}

func (v EventPostReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventPostIn struct {
	Name    string
	Company string
	AdminID string
}

type EventDeleteIn struct {
	ID string
}

type EventUpdateReq struct {
	Name    pkg.Varchar `json:"name"`
	Company pkg.Varchar `json:"company"`
}

func (v EventUpdateReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventUpdateIn struct {
	ID string
	EventUpdateReq
}
