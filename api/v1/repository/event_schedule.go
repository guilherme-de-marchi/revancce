package repository

import (
	"context"
	"fmt"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
)

func EventScheduleGet(ctx context.Context, in model.EventScheduleGetIn) ([]model.EventScheduleGetOut, error) {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("id", in.ID, "="),
			pkg.NewQueryParam("event", in.Event, "="),
			pkg.NewQueryParam("starts_at", in.From, ">="),
			pkg.NewQueryParam("ends_at", in.To, "<="),
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
					starts_at, 
					ends_at
				from events_schedules
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

	var out []model.EventScheduleGetOut
	for rows.Next() {
		var schedule model.EventScheduleGetOut
		err := rows.Scan(
			&schedule.ID,
			&schedule.Event,
			&schedule.StartsAt,
			&schedule.EndsAt,
		)
		if err != nil {
			return nil, pkg.Error(err)
		}
		out = append(out, schedule)
	}

	return out, nil
}

func EventSchedulePost(ctx context.Context, in model.EventSchedulePostIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			insert into events_schedules
			(event, starts_at, ends_at, created_by)
			values ($1, $2, $3, $4)
		`,
		in.Event,
		in.StartsAt,
		in.EndsAt,
		in.AdminID,
	)

	return pkg.Error(err)
}

func EventScheduleDelete(ctx context.Context, in model.EventScheduleDeleteIn) error {
	_, err := pkg.Database.Exec(
		ctx,
		`
			delete from events_schedules
			where id=$1
		`,
		in.ID,
	)

	return pkg.Error(err)
}

func EventScheduleUpdate(ctx context.Context, in model.EventScheduleUpdateIn) error {
	params, paramsValues := pkg.GenerateQueryParams(
		[]pkg.QueryParam{
			pkg.NewQueryParam("event", in.Event, ""),
			pkg.NewQueryParam("starts_at", in.StartsAt, ""),
			pkg.NewQueryParam("ends_at", in.EndsAt, ""),
		},
		"",
		",",
		2,
	)

	_, err := pkg.Database.Exec(
		ctx,
		fmt.Sprintf(
			`
			update events_schedules
			set %s
			where id=$1
			`,
			params,
		),
		append([]any{in.ID}, paramsValues...)...,
	)

	return pkg.Error(err)
}
