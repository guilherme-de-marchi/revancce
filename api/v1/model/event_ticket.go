package model

import (
	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

type EventTicketGetReq struct {
	ID     pkg.Varchar `form:"id"`
	Event  pkg.Varchar `form:"event"`
	Name   pkg.Varchar `form:"name"`
	Offset pkg.Integer `form:"offset"`
	Page   pkg.Integer `form:"page"`
	Limit  pkg.Integer `form:"limit"`
}

func (v EventTicketGetReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventTicketGetIn EventTicketGetReq

type EventTicketGetOut struct {
	ID    string `json:"id"`
	Event string `json:"event"`
	Name  string `json:"name"`
}

type EventTicketPostReq struct {
	Event pkg.Varchar `json:"event"`
	Name  pkg.Varchar `json:"name"`
}

func (v EventTicketPostReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventTicketPostIn struct {
	Event   string
	Name    string
	AdminID string
}

type EventTicketDeleteIn struct {
	ID string
}

type EventTicketUpdateReq struct {
	Event pkg.Varchar `json:"event"`
	Name  pkg.Varchar `json:"name"`
}

func (v EventTicketUpdateReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventTicketUpdateIn struct {
	ID string
	EventTicketUpdateReq
}
