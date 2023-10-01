package pkg

import "encoding/json"

type Integer struct {
	validate func(*int) error
	Value    *int
}

func (v *Integer) UnmarshalJSON(bytes []byte) error {
	var s int
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	v.Value = &s
	return nil
}

func (v Integer) Validate() error {
	if v.validate == nil {
		return validatorErrUnableToValidate
	}

	return v.validate(v.Value)
}

func NewInteger(require bool) Integer {
	return Integer{validate: func(v *int) error {
		if v == nil && require {
			return validatorErrRequired
		}

		return nil
	}}
}
