package model

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
