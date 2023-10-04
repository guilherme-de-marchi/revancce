package pkg

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	validatorErrUnableToValidate = errors.New("unable to validade field")
	validatorErrTooLarge         = errors.New("field is too large")
	validatorErrRequired         = errors.New("field is required")
)

type OptionalValue interface {
	Validate() error
	GetValue() any
	IsNil() bool
}

func ValidateStruct(s any) error {
	t := reflect.TypeOf(s)
	if t == nil || t.Kind() != reflect.Struct {
		return errors.New("target value is not a struct")
	}

	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		name := getFieldName(t.Field(i))

		if f.IsZero() {
			return fmt.Errorf("field '%s' is empty", name)
		}

		validator, ok := f.Interface().(OptionalValue)
		if !ok {
			continue
		}

		if err := validator.Validate(); err != nil {
			switch err {
			case validatorErrUnableToValidate:
				return fmt.Errorf("unable to validate field '%s'", name)
			case validatorErrTooLarge:
				return fmt.Errorf("field '%s' is too large", name)
			case validatorErrRequired:
				return fmt.Errorf("field '%s' is required", name)
			default:
				Log.Println(err)
				return fmt.Errorf("something went wrong validating field '%s'", name)
			}
		}
	}

	return nil
}

func getFieldName(f reflect.StructField) string {
	tags := strings.Split(f.Tag.Get("json"), ",")
	if len(tags) > 0 {
		return tags[0]
	}
	return f.Name
}
