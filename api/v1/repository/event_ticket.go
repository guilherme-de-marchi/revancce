package repository

import (
	"context"
	"fmt"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

func EventTicketGet(ctx context.Context, in model.EventTicketGetIn) ([]model.EventTicketGetOut, error) {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("id", in.ID, "="),
			pkg.NewQueryParam("event", in.Event, "="),
			pkg.NewQueryParam("name", in.Name, "="),
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
					event, 
					name
				from events_tickets
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

	var out []model.EventTicketGetOut
	for rows.Next() {
		var v model.EventTicketGetOut
		err := rows.Scan(
			&v.ID,
			&v.Event,
			&v.Name,
		)
		if err != nil {
			return nil, pkg.Error(err)
		}
		out = append(out, v)
	}

	return out, nil
}

func EventTicketPost(ctx context.Context, in model.EventTicketPostIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			insert into events_tickets
			(event, name, created_by)
			values ($1, $2, $3)
		`,
		in.Event,
		in.Name,
		in.AdminID,
	)

	return pkg.Error(err)
}

func EventTicketDelete(ctx context.Context, in model.EventTicketDeleteIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			delete from events_tickets
			where id=$1
		`,
		in.ID,
	)

	return pkg.Error(err)
}

func EventTicketUpdate(ctx context.Context, in model.EventTicketUpdateIn) error {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("event", in.Event, "="),
			pkg.NewQueryParam("name", in.Name, "="),
		},
		"",
		",",
		2,
	)

	_, err := pkg.Database.Exec(
		ctx,
		fmt.Sprintf(
			`
			update events_tickets
			set %s
			where id=$1
			`,
			params,
		),
		append([]any{in.ID}, paramsValues...)...,
	)

	return pkg.Error(err)
}
