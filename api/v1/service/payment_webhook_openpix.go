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

func PaymentWebhookOpenpix(ctx context.Context, in model.PaymentWebhookOpenpixPostIn) (int, any) {
	var err error
	switch in.Type {
	case "OPENPIX:CHARGE_COMPLETED":
		err = repository.PaymentWebhookOpenpixChargeCompleted(ctx, in)
	}
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
		pkg.Log.Println(err)
		err = errors.New("invalid value")
		status = http.StatusBadRequest
	case pgerrcode.UniqueViolation,
		pgerrcode.ForeignKeyViolation:
		err = errors.New("already exists")
		status = http.StatusUnauthorized
	default:
		pkg.Log.Println(err)
		err = errors.New("something went wrong")
		status = http.StatusInternalServerError
	}

	return status, pkg.ErrorMsg(err.Error())
}
