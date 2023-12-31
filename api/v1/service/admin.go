package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/repository"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(ctx context.Context, in model.AdminLoginIn) (int, any) {
	resp, err := repository.AdminLogin(ctx, in)
	if err == nil {
		return http.StatusOK, resp
	}

	var status int
	switch {
	case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword),
		errors.Is(err, pgx.ErrNoRows):
		err = errors.New("invalid credentials")
		status = http.StatusBadRequest
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}

func AdminRegister(ctx context.Context, in model.AdminRegisterIn) (int, any) {
	err := repository.AdminRegister(ctx, in)
	if err == nil {
		return http.StatusCreated, nil
	}

	pgErr := pkg.ErrorToPgError(err)
	if pgErr == nil {
		pkg.Log.Println(err)
		return http.StatusInternalServerError, pkg.ErrorMsg("something went wrong")
	}

	var status int
	switch pgErr.Code {
	case pgerrcode.UniqueViolation:
		err = errors.New("name in use")
		status = http.StatusBadRequest
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}
