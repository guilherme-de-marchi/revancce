package model

import (
	"errors"
	"time"
)

type AdminLoginReq struct {
	Name     *string `json:"name"`
	Password *string `json:"password"`
}

func (v AdminLoginReq) Validate() error {
	if v.Name == nil {
		return errors.New("field 'name' is empty")
	}

	if v.Password == nil {
		return errors.New("field 'password' is empty")
	}

	if len(*v.Name) > 20 {
		return errors.New("field 'name' too large")
	}

	if len(*v.Password) > 20 {
		return errors.New("field 'password' too large")
	}

	return nil
}

type AdminLoginIn struct {
	Name     string
	Password string
}

type AdminLoginOut struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type AdminRegisterReq struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type AdminRegisterIn struct {
	Name     string
	Email    string
	Password string
	ID       string
}

func (v AdminRegisterReq) Validate() error {
	if v.Name == nil {
		return errors.New("field 'name' is empty")
	}

	if v.Email == nil {
		return errors.New("field 'email' is empty")
	}

	if v.Password == nil {
		return errors.New("field 'password' is empty")
	}

	if len(*v.Name) > 20 {
		return errors.New("field 'name' too large")
	}

	if len(*v.Email) > 30 {
		return errors.New("field 'email' too large")
	}

	if len(*v.Password) > 20 {
		return errors.New("field 'password' too large")
	}

	return nil
}
