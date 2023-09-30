package model

import "errors"

type EventLocationGetReq struct {
	ID      *string `form:"id"`
	Event   *string `form:"event"`
	Country *string `form:"country"`
	State   *string `form:"state"`
	City    *string `form:"city"`
	Street  *string `form:"street"`
	Number  *string `form:"number"`
	Offset  *int    `form:"offset"`
	Page    *int    `form:"page"`
	Limit   *int    `form:"limit"`
}

type EventLocationGetIn EventLocationGetReq

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
	Event          *string `json:"event"`
	Country        *string `json:"country"`
	State          *string `json:"state"`
	City           *string `json:"city"`
	Street         *string `json:"street"`
	Number         *string `json:"number"`
	AdditionalInfo *string `json:"additional_info"`
	MapsURL        *string `json:"maps_url"`
}

func (v EventLocationPostReq) Validate() error {
	if v.Event == nil {
		return errors.New("field 'event' is empty")
	}

	if v.Country == nil {
		return errors.New("field 'country' is empty")
	}

	if v.State == nil {
		return errors.New("field 'state' is empty")
	}

	if v.City == nil {
		return errors.New("field 'city' is empty")
	}

	if v.Street == nil {
		return errors.New("field 'street' is empty")
	}

	if v.Number == nil {
		return errors.New("field 'number' is empty")
	}

	if v.AdditionalInfo == nil {
		return errors.New("field 'additional_info' is empty")
	}

	if v.MapsURL == nil {
		return errors.New("field 'maps_url' is empty")
	}

	if len(*v.Event) > 40 {
		return errors.New("field 'event' too large")
	}

	if len(*v.Country) > 20 {
		return errors.New("field 'country' too large")
	}

	if len(*v.State) > 20 {
		return errors.New("field 'state' too large")
	}

	if len(*v.City) > 20 {
		return errors.New("field 'city' too large")
	}

	if len(*v.Street) > 20 {
		return errors.New("field 'street' too large")
	}

	if len(*v.Number) > 20 {
		return errors.New("field 'number' too large")
	}

	if len(*v.AdditionalInfo) > 20 {
		return errors.New("field 'additional_info' too large")
	}

	if len(*v.MapsURL) > 20 {
		return errors.New("field 'maps_url' too large")
	}

	return nil
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
	Event          *string `json:"event"`
	Country        *string `json:"country"`
	State          *string `json:"state"`
	City           *string `json:"city"`
	Street         *string `json:"street"`
	Number         *string `json:"number"`
	AdditionalInfo *string `json:"additional_info"`
	MapsURL        *string `json:"maps_url"`
}

func (v EventLocationUpdateReq) Validate() error {
	if v.Event == nil &&
		v.Country == nil &&
		v.State == nil &&
		v.City == nil &&
		v.Street == nil &&
		v.Number == nil &&
		v.AdditionalInfo == nil &&
		v.MapsURL == nil {
		return errors.New("must provide at least one field")
	}

	return nil
}

type EventLocationUpdateIn struct {
	ID string
	EventLocationUpdateReq
}
