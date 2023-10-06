package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

var (
	ErrClientTicketPostLimitAmountReached = errors.New("limit amount of tickets reached")
	ErrClientTicketPostLimitTimeReached   = errors.New("limit time of this batch reached")
	ErrClientTicketPostClosed             = errors.New("batch closed")
	ErrClientTicketCheckinUsed            = errors.New("ticket has been used")
)

func ClientTicketGet(ctx context.Context, in model.ClientTicketGetIn) ([]model.ClientTicketGetOut, error) {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("id", in.ID, "="),
			pkg.NewQueryParam("client", in.Client, "="),
			pkg.NewQueryParam("batch", in.Batch, "="),
			pkg.NewQueryParam("transaction", in.Transaction, "="),
			pkg.NewQueryParam("created_at", in.From, ">="),
			pkg.NewQueryParam("created_at", in.To, "<="),
			pkg.NewQueryParam("used", in.Used, "="),
		},
		"where",
		"and",
		1,
	)

	if in.Limit.Value != nil {
		if *in.Limit.Value > 10 {
			*in.Limit.Value = 10
		} else if *in.Limit.Value == 0 {
			*in.Limit.Value = 1
		}
	} else {
		in.Limit.Value = pkg.Pointer(10)
	}

	paginations, paginationValues := pkg.GenerateQueryPagination(
		map[string]*int{
			"offset": pkg.CalcOptionalOffset(in.Offset.Value, in.Page.Value, in.Limit.Value),
			"limit":  in.Limit.Value,
		},
		len(paramsValues)+1,
	)

	rows, err := pkg.Database.Query(
		ctx,
		fmt.Sprintf(
			`
				select
					id,
					client,
					batch,
					transaction,
					used,
					created_at
				from clients_tickets
				%s 
				%s
			`,
			params,
			paginations,
		),
		append(paramsValues, paginationValues...)...,
	)
	if err != nil {
		return nil, pkg.Error(err)
	}

	if err := rows.Err(); err != nil {
		return nil, pkg.Error(err)
	}

	var out []model.ClientTicketGetOut
	for rows.Next() {
		var v model.ClientTicketGetOut
		err := rows.Scan(
			&v.ID,
			&v.Client,
			&v.Batch,
			&v.Transaction,
			&v.Used,
			&v.CreatedAt,
		)
		if err != nil {
			return nil, pkg.Error(err)
		}
		out = append(out, v)
	}

	return out, nil
}

func ClientTicketPost(ctx context.Context, in model.ClientTicketPostIn) (model.ClientTicketPostOut, error) {
	var out model.ClientTicketPostOut

	row := pkg.Database.QueryRow(
		ctx,
		`
			select count(*)
			from clients_tickets
			where batch=$1
		`,
		in.Batch,
	)

	var count int
	if err := row.Scan(&count); err != nil {
		return out, pkg.Error(err)
	}

	outEventBatch, err := EventBatchGet(
		ctx,
		model.EventBatchGetIn{ID: pkg.Varchar{Value: &in.Batch}},
	)
	if err != nil {
		return out, pkg.Error(err)
	}

	if len(outEventBatch) == 0 {
		return out, pkg.Error(ErrEventBatchNotFound)
	}

	if !outEventBatch[0].Opened {
		return out, pkg.Error(ErrClientTicketPostClosed)
	}

	if count >= outEventBatch[0].LimitAmount {
		EventBatchUpdate(ctx, model.EventBatchUpdateIn{
			ID: outEventBatch[0].ID,
			EventBatchUpdateReq: model.EventBatchUpdateReq{
				Opened: pkg.Boolean{Value: pkg.Pointer(false)},
			},
		})
		return out, pkg.Error(ErrClientTicketPostLimitAmountReached)
	}

	if time.Now().After(outEventBatch[0].LimitTime) {
		EventBatchUpdate(ctx, model.EventBatchUpdateIn{
			ID: outEventBatch[0].ID,
			EventBatchUpdateReq: model.EventBatchUpdateReq{
				Opened: pkg.Boolean{Value: pkg.Pointer(false)},
			},
		})
		return out, pkg.Error(ErrClientTicketPostLimitTimeReached)
	}

	row = pkg.Database.QueryRow(
		ctx,
		`
			insert into clients_tickets
			(client, batch, transaction, created_by)
			values ($1, $2, $3, $4)
			returning id
		`,
		in.Client,
		in.Batch,
		in.Transaction,
		in.AdminID,
	)

	return out, pkg.Error(row.Scan(&out.ID))
}

func ClientTicketDelete(ctx context.Context, in model.ClientTicketDeleteIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			delete from clients_tickets
			where id=$1
		`,
		in.ID,
	)

	return pkg.Error(err)
}

func ClientTicketUpdate(ctx context.Context, in model.ClientTicketUpdateIn) error {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("client", in.Client, "="),
			pkg.NewQueryParam("batch", in.Batch, "="),
			pkg.NewQueryParam("transaction", in.Transaction, "="),
			pkg.NewQueryParam("used", in.Used, "="),
		},
		"",
		",",
		2,
	)

	_, err := pkg.Database.Exec(
		ctx,
		fmt.Sprintf(
			`
			update clients_tickets
			set %s
			where id=$1
			`,
			params,
		),
		append([]any{in.ID}, paramsValues...)...,
	)

	return pkg.Error(err)
}

func ClientTicketCheckin(ctx context.Context, in model.ClientTicketCheckinIn) error {
	row := pkg.Database.QueryRow(
		ctx,
		`
			select used
			from clients_tickets
			where id=$1
		`,
		in.ID,
	)

	var used bool
	if err := row.Scan(&used); err != nil {
		return pkg.Error(err)
	}

	if used {
		return pkg.Error(ErrClientTicketCheckinUsed)
	}

	_, err := pkg.Database.Exec(
		ctx,
		`
			update clients_tickets
			set used=true
			where id=$1
		`,
		in.ID,
	)

	return pkg.Error(err)
}
