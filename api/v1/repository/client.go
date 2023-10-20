package repository

import (
	"context"
	"fmt"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

func ClientGet(ctx context.Context, in model.ClientGetIn) ([]model.ClientGetOut, error) {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("id", in.ID, "="),
			pkg.NewQueryParam("name", in.Name, "="),
			pkg.NewQueryParam("email", in.Email, "="),
			pkg.NewQueryParam("cpf", in.CPF, "="),
			pkg.NewQueryParam("phone", in.Phone, "="),
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
					email,
					cpf,
					phone
				from clients
				%s 
				%s
			`,
			params,
			paginations,
		),
		append(paramsValues, paginationValues...)...,
	)
	defer rows.Close()
	if err != nil {
		return nil, pkg.Error(err)
	}

	if err := rows.Err(); err != nil {
		return nil, pkg.Error(err)
	}

	var out []model.ClientGetOut
	for rows.Next() {
		var v model.ClientGetOut
		err := rows.Scan(
			&v.ID,
			&v.Name,
			&v.Email,
			&v.CPF,
			&v.Phone,
		)
		if err != nil {
			return nil, pkg.Error(err)
		}
		out = append(out, v)
	}

	return out, nil
}

func ClientPost(ctx context.Context, in model.ClientPostIn) (model.ClientPostOut, error) {
	row := pkg.Database.QueryRow(
		ctx,
		`
			insert into clients
			(name, email, cpf, phone, created_by)
			values ($1, $2, $3, $4, $5)
			returning id
		`,
		in.Name,
		in.Email,
		in.CPF,
		in.Phone,
		in.AdminID,
	)

	var out model.ClientPostOut
	return out, pkg.Error(row.Scan(&out.ID))
}

func ClientDelete(ctx context.Context, in model.ClientDeleteIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			delete from clients
			where id=$1
		`,
		in.ID,
	)

	return pkg.Error(err)
}

func ClientUpdate(ctx context.Context, in model.ClientUpdateIn) error {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("name", in.Name, "="),
			pkg.NewQueryParam("email", in.Email, "="),
			pkg.NewQueryParam("cpf", in.CPF, "="),
			pkg.NewQueryParam("phone", in.Phone, "="),
		},
		"",
		",",
		2,
	)

	_, err := pkg.Database.Exec(
		ctx,
		fmt.Sprintf(
			`
			update clients
			set %s
			where id=$1
			`,
			params,
		),
		append([]any{in.ID}, paramsValues...)...,
	)

	return pkg.Error(err)
}
