package repository

import (
	"context"
	"fmt"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

func EventBatchGet(ctx context.Context, in model.EventBatchGetIn) ([]model.EventBatchGetOut, error) {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("id", in.ID, "="),
			pkg.NewQueryParam("ticket", in.Ticket, "="),
			pkg.NewQueryParam("number", in.Number, "="),
			pkg.NewQueryParam("limit_amount", in.FromLimitAmount, ">="),
			pkg.NewQueryParam("limit_amount", in.ToLimitAmount, "<="),
			pkg.NewQueryParam("limit_time", in.FromLimitTime, ">="),
			pkg.NewQueryParam("limit_time", in.ToLimitTime, "<="),
			pkg.NewQueryParam("opened", in.Opened, "="),
			pkg.NewQueryParam("price", in.FromPrice, ">="),
			pkg.NewQueryParam("price", in.ToPrice, "<="),
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
					ticket, 
					number,
					limit_amount,
					limit_time,
					opened,
					price
				from events_batches
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

	var out []model.EventBatchGetOut
	for rows.Next() {
		var v model.EventBatchGetOut
		err := rows.Scan(
			&v.ID,
			&v.Ticket,
			&v.Number,
			&v.LimitAmount,
			&v.LimitTime,
			&v.Opened,
			&v.Price,
		)
		if err != nil {
			return nil, pkg.Error(err)
		}
		out = append(out, v)
	}

	return out, nil
}

func EventBatchPost(ctx context.Context, in model.EventBatchPostIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			insert into events_batches
			(ticket, limit_amount, limit_time, price, created_by)
			values ($1, $2, $3, $4, $5)
		`,
		in.Ticket,
		in.LimitAmount,
		in.LimitTime,
		in.Price,
		in.AdminID,
	)

	return pkg.Error(err)
}

func EventBatchDelete(ctx context.Context, in model.EventBatchDeleteIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			delete from events_batches
			where id=$1
		`,
		in.ID,
	)

	return pkg.Error(err)
}

func EventBatchUpdate(ctx context.Context, in model.EventBatchUpdateIn) error {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("ticket", in.Ticket, "="),
			pkg.NewQueryParam("limit_amount", in.LimitAmount, "="),
			pkg.NewQueryParam("limit_time", in.LimitTime, "="),
			pkg.NewQueryParam("opened", in.Opened, "="),
			pkg.NewQueryParam("price", in.Price, "="),
		},
		"",
		",",
		2,
	)

	_, err := pkg.Database.Exec(
		ctx,
		fmt.Sprintf(
			`
			update events_batches
			set %s
			where id=$1
			`,
			params,
		),
		append([]any{in.ID}, paramsValues...)...,
	)

	return pkg.Error(err)
}
