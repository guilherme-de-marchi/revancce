package model

import (
	"time"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

type EventBatchGetReq struct {
	ID              pkg.Varchar `form:"id"`
	Ticket          pkg.Varchar `form:"ticket"`
	Number          pkg.Integer `form:"number"`
	FromLimitAmount pkg.Integer `form:"from_limit_amount"`
	ToLimitAmount   pkg.Integer `form:"to_limit_amount"`
	FromLimitTime   pkg.Varchar `form:"from_limit_time"`
	ToLimitTime     pkg.Varchar `form:"to_limit_time"`
	Opened          pkg.Boolean `form:"opened"`
	FromPrice       pkg.Integer `form:"from_price"`
	ToPrice         pkg.Integer `form:"to_price"`
	Offset          pkg.Integer `form:"offset"`
	Page            pkg.Integer `form:"page"`
	Limit           pkg.Integer `form:"limit"`
}

func (v EventBatchGetReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventBatchGetIn EventBatchGetReq

type EventBatchGetOut struct {
	ID          string    `json:"id"`
	Ticket      string    `json:"ticket"`
	Number      int       `json:"number"`
	LimitAmount int       `json:"limit_amount"`
	LimitTime   time.Time `json:"limit_time"`
	Opened      bool      `json:"opened"`
	Price       int       `json:"price"`
}

type EventBatchPostReq struct {
	Ticket      pkg.Varchar `json:"ticket"`
	LimitAmount pkg.Integer `json:"limit_amount"`
	LimitTime   pkg.Varchar `json:"limit_time"`
	Price       pkg.Integer `json:"price"`
}

func (v EventBatchPostReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventBatchPostIn struct {
	Ticket      string
	Number      int
	LimitAmount int
	LimitTime   string
	Price       int
	AdminID     string
}

type EventBatchDeleteIn struct {
	ID string
}

type EventBatchUpdateReq struct {
	Ticket      pkg.Varchar `json:"ticket"`
	LimitAmount pkg.Integer `json:"limit_amount"`
	LimitTime   pkg.Varchar `json:"limit_time"`
	Opened      pkg.Boolean `json:"opened"`
	Price       pkg.Integer `json:"price"`
}

func (v EventBatchUpdateReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type EventBatchUpdateIn struct {
	ID string
	EventBatchUpdateReq
}
