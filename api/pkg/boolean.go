package pkg

import (
	"encoding/json"
)

type Boolean struct {
	validate func(*bool) error
	Value    *bool
}

func (v *Boolean) UnmarshalJSON(bytes []byte) error {
	var s bool
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	v.Value = &s
	return v.Validate()
}

func (v Boolean) Validate() error {
	if v.validate == nil {
		return validatorErrUnableToValidate
	}

	return v.validate(v.Value)
}

func (v Boolean) GetValue() any {
	return v.Value
}

func (v Boolean) IsNil() bool {
	return v.Value == nil
}

func NewBoolean(require bool) Boolean {
	return Boolean{validate: func(v *bool) error {
		if v == nil && require {
			return validatorErrRequired
		}

		return nil
	}}
}
