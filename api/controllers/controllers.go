package controllers

import (
	"reflect"
)

func Set(c any) {
	t := reflect.ValueOf(c)
	for i := 0; i < t.NumMethod(); i++ {
		t.Type().Method(i).Func.Call([]reflect.Value{t})
	}
}
