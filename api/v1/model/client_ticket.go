package model

import (
	"time"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

type ClientTicketGetReq struct {
	ID          pkg.Varchar `form:"id"`
	Client      pkg.Varchar `form:"client"`
	Batch       pkg.Varchar `form:"batch"`
	Transaction pkg.Varchar `form:"transaction"`
	From        pkg.Varchar `form:"from"`
	To          pkg.Varchar `form:"to"`
	Used        pkg.Boolean `form:"used"`
	Offset      pkg.Integer `form:"offset"`
	Page        pkg.Integer `form:"page"`
	Limit       pkg.Integer `form:"limit"`
}

func (v ClientTicketGetReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type ClientTicketGetIn ClientTicketGetReq

type ClientTicketGetOut struct {
	ID          string    `json:"id"`
	Client      string    `json:"client"`
	Batch       string    `json:"batch"`
	Transaction string    `json:"transaction"`
	Used        bool      `json:"used"`
	CreatedAt   time.Time `json:"created_at"`
}

type ClientTicketPostReq struct {
	Client      pkg.Varchar `json:"client"`
	Batch       pkg.Varchar `json:"batch"`
	Transaction pkg.Varchar `json:"transaction"`
}

func (v ClientTicketPostReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type ClientTicketPostIn struct {
	Client      string
	Batch       string
	Transaction string
	AdminID     string
}

type ClientTicketPostOut struct {
	ID string
}

type ClientTicketDeleteIn struct {
	ID string
}

type ClientTicketUpdateReq struct {
	Client      pkg.Varchar `json:"client"`
	Batch       pkg.Varchar `json:"batch"`
	Transaction pkg.Varchar `json:"transaction"`
	Used        pkg.Boolean `json:"used"`
}

func (v ClientTicketUpdateReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type ClientTicketUpdateIn struct {
	ID string
	ClientTicketUpdateReq
}

type ClientTicketCheckinReq struct {
	ID pkg.Varchar `json:"id"`
}

func (v ClientTicketCheckinReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type ClientTicketCheckinIn struct {
	ID string
}
