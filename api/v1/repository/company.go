package repository

import (
	"context"
	"fmt"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

func CompanyGet(ctx context.Context, in model.CompanyGetIn) ([]model.CompanyGetOut, error) {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("id", in.ID, "="),
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
					name
				from companies
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

	var out []model.CompanyGetOut
	for rows.Next() {
		var v model.CompanyGetOut
		err := rows.Scan(
			&v.ID,
			&v.Name,
		)
		if err != nil {
			return nil, pkg.Error(err)
		}
		out = append(out, v)
	}

	return out, nil
}

func CompanyPost(ctx context.Context, in model.CompanyPostIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			insert into companies
			(name, created_by)
			values ($1, $2)
		`,
		in.Name,
		in.AdminID,
	)

	return pkg.Error(err)
}

func CompanyDelete(ctx context.Context, in model.CompanyDeleteIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			delete from companies
			where id=$1
		`,
		in.ID,
	)

	return pkg.Error(err)
}

func CompanyUpdate(ctx context.Context, in model.CompanyUpdateIn) error {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
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
			update companies
			set %s
			where id=$1
			`,
			params,
		),
		append([]any{in.ID}, paramsValues...)...,
	)

	return pkg.Error(err)
}
