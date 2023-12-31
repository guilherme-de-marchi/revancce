package pkg

import (
	"testing"
)

func TestGenerateQueryConditionals(t *testing.T) {
	a := NewVarchar(3, false)
	a.Value = Pointer("def")

	b := NewVarchar(3, false)
	b.Value = Pointer("jkl")

	query, values := GenerateQueryParams(
		[]QueryParam{
			NewQueryParam("abc", a, "="),
			NewQueryParam("ghi", a, "="),
		},
		"where",
		"and",
		1,
	)
	if query != "where abc = $1 and ghi = $2 " {
		t.Error("invalid query")
		t.Log(query)
	}
	if len(values) != 2 {
		t.Error("invalid values")
		t.Log(query, values)
	}

	query, values = GenerateQueryParams(
		[]QueryParam{
			NewQueryParam("abc", a, "="),
			NewQueryParam("ghi", a, "="),
		},
		"set",
		",",
		2,
	)
	if query != "set abc = $2 , ghi = $3 " {
		t.Error("invalid query")
		t.Log(query)
	}
	if len(values) != 2 {
		t.Error("invalid values")
		t.Log(query, values)
	}

	query, values = GenerateQueryParams(
		[]QueryParam{},
		"where",
		"and",
		0,
	)
	if query != "" {
		t.Error("invalid query, should be empty")
		t.Log(query)
	}
	if len(values) != 0 {
		t.Error("invalid values, should be empty")
		t.Log(query, values)
	}
}

func TestGenerateQueryPagination(t *testing.T) {
	query, values := GenerateQueryPagination(
		map[string]*int{
			"abc": Pointer(1),
			"ghi": Pointer(2),
		},
		1,
	)
	if query != "abc $1 ghi $2 " {
		t.Error("invalid query")
		t.Log(query)
	}
	if len(values) != 2 {
		t.Error("invalid values")
		t.Log(query, values)
	}

	query, values = GenerateQueryPagination(map[string]*int{}, 0)
	if query != "" {
		t.Error("invalid query, should be empty")
		t.Log(query)
	}
	if len(values) != 0 {
		t.Error("invalid values, should be empty")
		t.Log(query, values)
	}
}
