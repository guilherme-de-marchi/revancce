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

type ClientTicketPurchasePostReq struct {
	Batch pkg.Varchar `json:"batch"`
}

func (v ClientTicketPurchasePostReq) Validate() error {
	return pkg.ValidateStruct(v)
}

type ClientTicketPurchasePostIn struct {
	Batch string
}

type ClientTicketPurchasePostOut struct {
	PaymentLinkURL string
}

type OpenpixCreateChargeReq struct {
	CorrelationID  string                  `json:"correlationID"`
	Value          int                     `json:"value"`
	Type           string                  `json:"type"`
	Comment        string                  `json:"comment"`
	ExpiresIn      int                     `json:"expiresIn"`
	AdditionalInfo []OpenpixAdditionalInfo `json:"additionalInfo"`
}

type OpenpixAdditionalInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OpenpixCreateChargeResp struct {
	Charge struct {
		Customer       interface{} `json:"customer"`
		Value          int         `json:"value"`
		Comment        string      `json:"comment"`
		Identifier     string      `json:"identifier"`
		CorrelationID  string      `json:"correlationID"`
		PaymentLinkID  string      `json:"paymentLinkID"`
		TransactionID  string      `json:"transactionID"`
		Status         string      `json:"status"`
		AdditionalInfo []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"additionalInfo"`
		Discount          int       `json:"discount"`
		ValueWithDiscount int       `json:"valueWithDiscount"`
		ExpiresDate       time.Time `json:"expiresDate"`
		Type              string    `json:"type"`
		CreatedAt         time.Time `json:"createdAt"`
		UpdatedAt         time.Time `json:"updatedAt"`
		BrCode            string    `json:"brCode"`
		ExpiresIn         int       `json:"expiresIn"`
		PixKey            string    `json:"pixKey"`
		PaymentLinkURL    string    `json:"paymentLinkUrl"`
		QrCodeImage       string    `json:"qrCodeImage"`
		GlobalID          string    `json:"globalID"`
	} `json:"charge"`
	CorrelationID string `json:"correlationID"`
	BrCode        string `json:"brCode"`
}
