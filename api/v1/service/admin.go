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
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(ctx context.Context, req model.AdminLoginReq) (int, any) {
	resp, err := repository.AdminLogin(ctx, req)
	if err == nil {
		return http.StatusOK, resp
	}

	var status int
	switch {
	case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword),
		errors.Is(err, pgx.ErrNoRows):
		err = errors.New("invalid credentials")
		status = http.StatusUnauthorized
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}

func AdminRegister(ctx context.Context, req model.AdminRegisterReq) (int, any) {
	err := repository.AdminRegister(ctx, req)
	if err == nil {
		return http.StatusCreated, nil
	}

	pkgErr, ok := err.(pkg.Err)
	if !ok {
		pkg.Log.Println(err)
		return http.StatusInternalServerError, errors.New("something went wrong")
	}

	pgErr, ok := pkgErr.Err.(*pgconn.PgError)
	if !ok {
		pkg.Log.Println(err)
		return http.StatusInternalServerError, errors.New("something went wrong")
	}

	var status int
	switch pgErr.Code {
	case pgerrcode.UniqueViolation:
		err = errors.New("name in use")
		status = http.StatusUnauthorized
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}
