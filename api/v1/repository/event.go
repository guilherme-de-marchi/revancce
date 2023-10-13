package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

var ErrEventNotFound = errors.New("event not found")

func EventGet(ctx context.Context, in model.EventGetIn) ([]model.EventGetOut, error) {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("id", in.ID, "="),
			pkg.NewQueryParam("name", in.Name, "="),
			pkg.NewQueryParam("company", in.Company, "="),
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
					name,
					company
				from events
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

	var out []model.EventGetOut
	for rows.Next() {
		var v model.EventGetOut
		if err := rows.Scan(&v.ID, &v.Name, &v.Company); err != nil {
			return nil, pkg.Error(err)
		}
		out = append(out, v)
	}

	return out, nil
}

func EventPost(ctx context.Context, in model.EventPostIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			insert into events
			(name, company, created_by)
			values ($1, $2, $3)
		`,
		in.Name, in.Company, in.AdminID,
	)

	return pkg.Error(err)
}

func EventDelete(ctx context.Context, in model.EventDeleteIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			delete from events
			where id=$1
		`,
		in.ID,
	)

	return pkg.Error(err)
}

func EventUpdate(ctx context.Context, in model.EventUpdateIn) error {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("name", in.Name, "="),
			pkg.NewQueryParam("company", in.Company, "="),
		},
		"",
		",",
		2,
	)

	_, err := pkg.Database.Exec(
		ctx,
		fmt.Sprintf(
			`
			update events
			set %s
			where id=$1
			`,
			params,
		),
		append([]any{in.ID}, paramsValues...)...,
	)

	return pkg.Error(err)
}
