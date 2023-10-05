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

func EventTicketGet(ctx context.Context, in model.EventTicketGetIn) (int, any) {
	resp, err := repository.EventTicketGet(ctx, in)
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

func EventTicketPost(ctx context.Context, in model.EventTicketPostIn) (int, any) {
	err := repository.EventTicketPost(ctx, in)
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
		pgerrcode.ForeignKeyViolation,
		pgerrcode.InvalidTextRepresentation:
		err = errors.New("invalid value")
		status = http.StatusBadRequest
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}

func EventTicketDelete(ctx context.Context, in model.EventTicketDeleteIn) (int, any) {
	err := repository.EventTicketDelete(ctx, in)
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
		status = http.StatusBadRequest
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}

func EventTicketUpdate(ctx context.Context, in model.EventTicketUpdateIn) (int, any) {
	err := repository.EventTicketUpdate(ctx, in)
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
		pgerrcode.InvalidTextRepresentation,
		pgerrcode.UniqueViolation:
		err = errors.New("invalid field")
		status = http.StatusBadRequest
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}
