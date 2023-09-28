package model

import "errors"

type EventGetReq struct {
	ID      *string `form:"id"`
	Name    *string `form:"name"`
	Company *string `form:"company"`
	Offset  *int    `form:"offset"`
	Page    *int    `form:"page"`
	Limit   *int    `form:"limit"`
}

type EventGetIn EventGetReq

type EventGetOut struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
}

type EventPostReq struct {
	Name    *string `json:"name"`
	Company *string `json:"company"`
}

func (v EventPostReq) Validate() error {
	if v.Name == nil {
		return errors.New("field 'name' is empty")
	}

	if v.Company == nil {
		return errors.New("field 'company' is empty")
	}

	if len(*v.Name) > 20 {
		return errors.New("field 'name' too large")
	}

	if len(*v.Company) > 40 {
		return errors.New("field 'company' too large")
	}

	return nil
}

type EventPostIn struct {
	Name    string
	Company string
	AdminID string
}

type EventDeleteIn struct {
	ID      string
	AdminID string
}
