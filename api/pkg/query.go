package pkg

import "fmt"

type QueryParam struct {
	Key       string
	Value     OptionalValue
	Condition string
}

func NewQueryParam(key string, value OptionalValue, condition string) QueryParam {
	return QueryParam{
		Key:       key,
		Value:     value,
		Condition: condition,
	}
}

func GenerateQueryParams(params []QueryParam, init, sep string, countStart int) (string, []any) {
	if len(params) == 0 {
		return "", nil
	}

	var values []any
	var query string
	count := countStart
	for _, p := range params {
		if p.Value.IsNil() {
			continue
		}

		if count == countStart {
			query += init
		} else {
			query += sep
		}

		query += fmt.Sprintf(" %s %s $%v ", p.Key, p.Condition, count)
		values = append(values, p.Value.GetValue())
		count++
	}

	return query, values
}

func GenerateQueryPagination[T any](paginations map[string]*T, countStart int) (string, []any) {
	if len(paginations) == 0 {
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
