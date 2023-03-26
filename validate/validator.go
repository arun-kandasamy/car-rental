package validate

import policy "github.com/example/car-rental-service/proto/generated/policy"

type validator interface {
	validate(*field)
	setNext(validator)
}

type field struct {
	validationError error
	name            string
	value           interface{}
}

func newField(name string, value interface{}) *field {
	f := &field{
		name:  name,
		value: value,
	}

	return f
}

func (f *field) runValidator(options *policy.IdValidationOpts) error {
	rv := &requiredValidator{options: options}
	lv := &lengthValidator{options: options}
	tv := &typeValidator{options: options}

	// wire the validators
	// required --> type --> length
	rv.setNext(tv)
	tv.setNext(lv)

	rv.validate(f)

	return f.validationError
}
