package repository

import (
	"context"
	"fmt"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

func EventGet(ctx context.Context, in model.EventGetIn) ([]model.EventGetOut, error) {
	params, paramsValues := pkg.GenerateQueryParams(
		map[string]*string{
			"id":      in.ID,
			"name":    in.Name,
			"company": in.Company,
		},
		"where",
		"and",
		1,
	)

	if in.Limit != nil {
		if *in.Limit > 10 {
			*in.Limit = 10
		} else if *in.Limit == 0 {
			*in.Limit = 1
		}
	} else {
		in.Limit = pkg.Pointer(10)
	}

	paginations, paginationValues := pkg.GenerateQueryPagination(
		map[string]*int{
			"offset": pkg.CalcOptionalOffset(in.Offset, in.Page, in.Limit),
			"limit":  in.Limit,
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
		var event model.EventGetOut
		if err := rows.Scan(&event.ID, &event.Name, &event.Company); err != nil {
			return nil, pkg.Error(err)
		}
		out = append(out, event)
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
		map[string]*string{
			"name":    in.Name,
			"company": in.Company,
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
