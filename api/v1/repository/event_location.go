package repository

import (
	"context"
	"fmt"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

func EventLocationGet(ctx context.Context, in model.EventLocationGetIn) ([]model.EventLocationGetOut, error) {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("id", in.ID, "="),
			pkg.NewQueryParam("event", in.Event, "="),
			pkg.NewQueryParam("country", in.Country, "="),
			pkg.NewQueryParam("state", in.State, "="),
			pkg.NewQueryParam("city", in.City, "="),
			pkg.NewQueryParam("street", in.Street, "="),
			pkg.NewQueryParam("number", in.Number, "="),
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
					country, 
					state, 
					city, 
					street, 
					number, 
					additional_info, 
					maps_url
				from events_locations
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

	var out []model.EventLocationGetOut
	for rows.Next() {
		var location model.EventLocationGetOut
		err := rows.Scan(
			&location.ID,
			&location.Event,
			&location.Country,
			&location.State,
			&location.City,
			&location.Street,
			&location.Number,
			&location.AdditionalInfo,
			&location.MapsURL,
		)
		if err != nil {
			return nil, pkg.Error(err)
		}
		out = append(out, location)
	}

	return out, nil
}

func EventLocationPost(ctx context.Context, in model.EventLocationPostIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			insert into events_locations
			(event, country, state, city, street, number, additional_info, maps_url, created_by)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`,
		in.Event,
		in.Country,
		in.State,
		in.City,
		in.Street,
		in.Number,
		in.AdditionalInfo,
		in.MapsURL,
		in.AdminID,
	)

	return pkg.Error(err)
}

func EventLocationDelete(ctx context.Context, in model.EventLocationDeleteIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			delete from events_locations
			where id=$1
		`,
		in.ID,
	)

	return pkg.Error(err)
}

func EventLocationUpdate(ctx context.Context, in model.EventLocationUpdateIn) error {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("event", in.Event, "="),
			pkg.NewQueryParam("country", in.Country, "="),
			pkg.NewQueryParam("state", in.State, "="),
			pkg.NewQueryParam("city", in.City, "="),
			pkg.NewQueryParam("street", in.Street, "="),
			pkg.NewQueryParam("number", in.Number, "="),
			pkg.NewQueryParam("additional_info", in.AdditionalInfo, "="),
			pkg.NewQueryParam("maps_url", in.MapsURL, "="),
		},
		"",
		",",
		2,
	)

	_, err := pkg.Database.Exec(
		ctx,
		fmt.Sprintf(
			`
			update events_locations
			set %s
			where id=$1
			`,
			params,
		),
		append([]any{in.ID}, paramsValues...)...,
	)

	return pkg.Error(err)
}
