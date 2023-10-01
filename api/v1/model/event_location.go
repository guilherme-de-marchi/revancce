package model

import (
	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

type EventLocationGetReq struct {
	ID      pkg.Varchar `form:"id"`
	Event   pkg.Varchar `form:"event"`
	Country pkg.Varchar `form:"country"`
	State   pkg.Varchar `form:"state"`
	City    pkg.Varchar `form:"city"`
	Street  pkg.Varchar `form:"street"`
	Number  pkg.Varchar `form:"number"`
	Offset  pkg.Integer `form:"offset"`
	Page    pkg.Integer `form:"page"`
	Limit   pkg.Integer `form:"limit"`
}

type EventLocationGetIn struct {
	ID      *string
	Event   *string
	Country *string
	State   *string
	City    *string
	Street  *string
	Number  *string
	Offset  *int
	Page    *int
	Limit   *int
}

type EventLocationGetOut struct {
	ID             string `json:"id"`
	Event          string `json:"event"`
	Country        string `json:"country"`
	State          string `json:"state"`
	City           string `json:"city"`
	Street         string `json:"street"`
	Number         string `json:"number"`
	AdditionalInfo string `json:"additional_info"`
	MapsURL        string `json:"maps_url"`
}

type EventLocationPostReq struct {
	Event          pkg.Varchar `json:"event"`
	Country        pkg.Varchar `json:"country"`
	State          pkg.Varchar `json:"state"`
	City           pkg.Varchar `json:"city"`
	Street         pkg.Varchar `json:"street"`
	Number         pkg.Varchar `json:"number"`
	AdditionalInfo pkg.Varchar `json:"additional_info"`
	MapsURL        pkg.Varchar `json:"maps_url"`
}

func (v EventLocationPostReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventLocationPostIn struct {
	Event          string
	Country        string
	State          string
	City           string
	Street         string
	Number         string
	AdditionalInfo string
	MapsURL        string
	AdminID        string
}

type EventLocationDeleteIn struct {
	ID string
}

type EventLocationUpdateReq struct {
	Event          pkg.Varchar `json:"event"`
	Country        pkg.Varchar `json:"country"`
	State          pkg.Varchar `json:"state"`
	City           pkg.Varchar `json:"city"`
	Street         pkg.Varchar `json:"street"`
	Number         pkg.Varchar `json:"number"`
	AdditionalInfo pkg.Varchar `json:"additional_info"`
	MapsURL        pkg.Varchar `json:"maps_url"`
}

func (v EventLocationUpdateReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventLocationUpdateIn struct {
	ID             string
	Event          *string
	Country        *string
	State          *string
	City           *string
	Street         *string
	Number         *string
	AdditionalInfo *string
	MapsURL        *string
}
