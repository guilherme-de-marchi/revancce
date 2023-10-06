package model

import (
	"errors"
	"strconv"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

type ClientGetReq struct {
	ID     pkg.Varchar `form:"id"`
	Name   pkg.Varchar `form:"name"`
	Email  pkg.Varchar `form:"email"`
	CPF    pkg.Varchar `form:"cpf"`
	Phone  pkg.Varchar `form:"phone"`
	Offset pkg.Integer `form:"offset"`
	Page   pkg.Integer `form:"page"`
	Limit  pkg.Integer `form:"limit"`
}

func (v ClientGetReq) Validate() error {
	if err := pkg.ValidateStruct(v); err != nil {
		return err
	}

	if v.CPF.Value != nil {
		if _, err := strconv.Atoi(*v.CPF.Value); err != nil {
			return errors.New("invalid cpf")
		}
	}

	if v.Phone.Value != nil {
		if _, err := strconv.Atoi(*v.Phone.Value); err != nil {
			return errors.New("invalid phone")
		}
	}

	return nil
}

type ClientGetIn ClientGetReq

type ClientGetOut struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
	Phone string `json:"phone"`
}

type ClientPostReq struct {
	Name  pkg.Varchar `json:"name"`
	Email pkg.Varchar `json:"email"`
	CPF   pkg.Varchar `json:"cpf"`
	Phone pkg.Varchar `json:"phone"`
}

func (v ClientPostReq) Validate() error {
	if err := pkg.ValidateStruct(v); err != nil {
		return err
	}

	if v.CPF.Value != nil {
		if _, err := strconv.Atoi(*v.CPF.Value); err != nil {
			return errors.New("invalid cpf")
		}
	}

	if v.Phone.Value != nil {
		if _, err := strconv.Atoi(*v.Phone.Value); err != nil {
			return errors.New("invalid phone")
		}
	}

	return nil
}

type ClientPostIn struct {
	Name    string
	Email   string
	CPF     string
	Phone   string
	AdminID string
}

type ClientPostOut struct {
	ID string `json:"id"`
}

type ClientDeleteIn struct {
	ID string
}

type ClientUpdateReq struct {
	Name  pkg.Varchar `json:"name"`
	Email pkg.Varchar `json:"email"`
	CPF   pkg.Varchar `json:"cpf"`
	Phone pkg.Varchar `json:"phone"`
}

func (v ClientUpdateReq) Validate() error {
	if err := pkg.ValidateStruct(v); err != nil {
		return err
	}

	if v.CPF.Value != nil {
		if _, err := strconv.Atoi(*v.CPF.Value); err != nil {
			return errors.New("invalid cpf")
		}
	}

	if v.Phone.Value != nil {
		if _, err := strconv.Atoi(*v.Phone.Value); err != nil {
			return errors.New("invalid phone")
		}
	}

	return nil
}

type ClientUpdateIn struct {
	ID string
	ClientUpdateReq
}
