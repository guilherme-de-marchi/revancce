package pkg

import (
	"fmt"
)

func GenerateQueryParams[T any](params map[string]*T, init, sep string, countStart int) (string, []any) {
	if len(params) == 0 {
		return "", nil
	}

	var values []any
	var query string
	count := countStart
	for k, v := range params {
		if v == nil {
			continue
		}

		if count == countStart {
			query += init
		} else {
			query += sep
		}

		query += fmt.Sprintf(" %s=$%v ", k, count)
		values = append(values, *v)
		count++
	}

	return query, values
}

func GenerateQueryPagination[T any](paginations map[string]*T, countStart int) (string, []any) {
	if len(paginations) == countStart {
		return "", nil
	}

	var values []any
	var query string
	count := countStart
	for k, v := range paginations {
		if v == nil {
			continue
		}

		query += fmt.Sprintf("%s $%v ", k, count)
		values = append(values, *v)
		count++
	}

	return query, values
}

func CalcOptionalOffset(optOffset, optPage, optLimit *int) *int {
	var offset, page, limit int

	if optOffset != nil {
		offset = *optOffset
	}
	if optPage != nil {
		page = *optPage
	}
	if optLimit != nil {
		limit = *optLimit
	}

	newOffset := page*limit + offset
	if newOffset == 0 {
		return nil
	}
	return &newOffset
}
