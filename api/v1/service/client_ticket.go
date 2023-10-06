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
)

func ClientTicketGet(ctx context.Context, in model.ClientTicketGetIn) (int, any) {
	resp, err := repository.ClientTicketGet(ctx, in)
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

func ClientTicketPost(ctx context.Context, in model.ClientTicketPostIn) (int, any) {
	_, err := repository.ClientTicketPost(ctx, in)
	if err == nil {
		return http.StatusCreated, nil
	}

	if errors.Is(err, repository.ErrClientTicketPostClosed) ||
		errors.Is(err, repository.ErrClientTicketPostLimitAmountReached) ||
		errors.Is(err, repository.ErrClientTicketPostLimitTimeReached) {
		return http.StatusUnauthorized, pkg.ErrorMsg("unable to create ticket")
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

func ClientTicketDelete(ctx context.Context, in model.ClientTicketDeleteIn) (int, any) {
	err := repository.ClientTicketDelete(ctx, in)
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

func ClientTicketUpdate(ctx context.Context, in model.ClientTicketUpdateIn) (int, any) {
	err := repository.ClientTicketUpdate(ctx, in)
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

func ClientTicketCheckin(ctx context.Context, in model.ClientTicketCheckinIn) (int, any) {
	err := repository.ClientTicketCheckin(ctx, in)
	if err == nil {
		return http.StatusOK, nil
	}

	if errors.Is(err, repository.ErrClientTicketCheckinUsed) {
		return http.StatusUnauthorized, pkg.ErrorMsg("ticket already used")
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return http.StatusUnauthorized, pkg.ErrorMsg("ticket not found")
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
