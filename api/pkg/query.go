package pkg

import "fmt"

func GenerateQueryConditionals(conditions map[string]*string, operator string, countStart int) (string, []any) {
	if len(conditions) == 0 {
		return "", nil
	}

	var values []any
	var query string
	count := countStart + 1
	for k, v := range conditions {
		if v == nil {
			continue
		}
		if count == 1 {
			query += "where"
		} else {
			query += operator
		}
		query += fmt.Sprintf(" %s=$%v ", k, count)
		values = append(values, *v)
		count++
	}

	return query, values
}

func GenerateQueryPagination(paginations map[string]*int, countStart int) (string, []any) {
	if len(paginations) == 0 {
		return "", nil
	}

	var values []any
	var query string
	count := countStart + 1
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
