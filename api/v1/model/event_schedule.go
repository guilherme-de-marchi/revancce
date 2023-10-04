package model

import (
	"time"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

type EventScheduleGetReq struct {
	ID     pkg.Varchar `form:"id"`
	Event  pkg.Varchar `form:"event"`
	From   pkg.Varchar `form:"from"`
	To     pkg.Varchar `form:"to"`
	Offset pkg.Integer `form:"offset"`
	Page   pkg.Integer `form:"page"`
	Limit  pkg.Integer `form:"limit"`
}

type EventScheduleGetIn EventScheduleGetReq

type EventScheduleGetOut struct {
	ID       string    `json:"id"`
	Event    string    `json:"event"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
}

type EventSchedulePostReq struct {
	Event    pkg.Varchar `json:"event"`
	StartsAt pkg.Varchar `json:"starts_at"`
	EndsAt   pkg.Varchar `json:"ends_at"`
}

func (v EventSchedulePostReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventSchedulePostIn struct {
	Event    string
	StartsAt string
	EndsAt   string
	AdminID  string
}

type EventScheduleDeleteIn struct {
	ID string
}

type EventScheduleUpdateReq struct {
	Event    pkg.Varchar `json:"event"`
	StartsAt pkg.Varchar `json:"starts_at"`
	EndsAt   pkg.Varchar `json:"ends_at"`
}

func (v EventScheduleUpdateReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventScheduleUpdateIn struct {
	ID string
	EventScheduleUpdateReq
}
