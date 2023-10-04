package pkg

import (
	"errors"
	"reflect"
	"strconv"
)

func BindQuery[T any](m map[string]string, obj *T) error {
	if obj == nil {
		return errors.New("unable to bind nil object")
	}

	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(*obj)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Tag.Get("form")
		vv, ok := m[name]
		if !ok {
			continue
		}

		x := v.Elem().Field(i).Addr().Interface()

		n, err := strconv.Atoi(vv)
		if err == nil {
			xx := x.(*Integer)
			xx.Value = &n
			continue
		}

		xx := x.(*Varchar)
		xx.Value = &vv
	}

	return nil
}
