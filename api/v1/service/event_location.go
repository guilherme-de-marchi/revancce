package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/repository"
	"github.com/jackc/pgerrcode"
)

func EventLocationGet(ctx context.Context, in model.EventLocationGetIn) (int, any) {
	resp, err := repository.EventLocationGet(ctx, in)
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

func EventLocationPost(ctx context.Context, in model.EventLocationPostIn) (int, any) {
	err := repository.EventLocationPost(ctx, in)
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
	case pgerrcode.UniqueViolation,
		pgerrcode.ForeignKeyViolation:
		err = errors.New("invalid value")
		status = http.StatusUnauthorized
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}

func EventLocationDelete(ctx context.Context, in model.EventLocationDeleteIn) (int, any) {
	err := repository.EventLocationDelete(ctx, in)
	if err == nil {
		return http.StatusOK, nil
	}

	pgErr := pkg.ErrorToPgError(err)
	if pgErr == nil {
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

func EventLocationUpdate(ctx context.Context, in model.EventLocationUpdateIn) (int, any) {
	err := repository.EventLocationUpdate(ctx, in)
	if err == nil {
		return http.StatusOK, nil
	}

	pgErr := pkg.ErrorToPgError(err)
	if pgErr == nil {
		pkg.Log.Println(err)
		return http.StatusInternalServerError, pkg.ErrorMsg("something went wrong")
	}

	var status int
	switch pgErr.Code {
	case pgerrcode.ForeignKeyViolation,
		pgerrcode.InvalidTextRepresentation:
		err = errors.New("invalid field")
		status = http.StatusUnauthorized
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}
