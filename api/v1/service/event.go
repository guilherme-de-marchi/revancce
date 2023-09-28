package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/repository"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func EventGet(ctx context.Context, in model.EventGetIn) (int, any) {
	resp, err := repository.EventGet(ctx, in)
	if err == nil {
		if len(resp) == 0 {
			return http.StatusNoContent, nil
		}
		return http.StatusOK, resp
	}

	pkg.Log.Println(err)
	err = errors.New("something went wrong")
	return http.StatusInternalServerError, pkg.ErrorMsg(err.Error())
}

func EventPost(ctx context.Context, in model.EventPostIn) (int, any) {
	err := repository.EventPost(ctx, in)
	if err == nil {
		return http.StatusCreated, nil
	}

	pkgErr, ok := err.(pkg.Err)
	if !ok {
		pkg.Log.Println(err)
		return http.StatusInternalServerError, pkg.ErrorMsg("something went wrong")
	}

	pgErr, ok := pkgErr.Err.(*pgconn.PgError)
	if !ok {
		pkg.Log.Println(err)
		return http.StatusInternalServerError, pkg.ErrorMsg("something went wrong")
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

func EventDelete(ctx context.Context, in model.EventDeleteIn) (int, any) {
	err := repository.EventDelete(ctx, in)
	if err == nil {
		return http.StatusOK, nil
	}

	pkgErr, ok := err.(pkg.Err)
	if !ok {
		pkg.Log.Println(err)
		return http.StatusInternalServerError, pkg.ErrorMsg("something went wrong")
	}

	pgErr, ok := pkgErr.Err.(*pgconn.PgError)
	if !ok {
		pkg.Log.Println(err)
		return http.StatusInternalServerError, pkg.ErrorMsg("something went wrong")
	}

	var status int
	switch pgErr.Code {
	case pgerrcode.InvalidTextRepresentation:
		err = errors.New("invalid field")
		status = http.StatusUnauthorized
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}
