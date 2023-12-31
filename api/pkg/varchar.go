package pkg

import (
	"encoding/json"
)

type Varchar struct {
	validate func(*string) error
	Value    *string
}

func (v *Varchar) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	v.Value = &s
	return v.Validate()
}

func (v Varchar) Validate() error {
	if v.validate == nil {
		return validatorErrUnableToValidate
	}

	return v.validate(v.Value)
}

func (v Varchar) GetValue() any {
	return v.Value
}

func (v Varchar) IsNil() bool {
	return v.Value == nil
}

func NewVarchar(size int, require bool) Varchar {
	return Varchar{validate: func(v *string) error {
		if v == nil {
			if require {
				return validatorErrRequired
			}

			return nil
		}

		if len(*v) > size {
			return validatorErrTooLarge
		}
		return nil
	}}
}
