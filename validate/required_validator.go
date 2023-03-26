package validate

import (
	"fmt"
	"strings"

	policy "github.com/example/car-rental-service/proto/generated/policy"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type requiredValidator struct {
	options *policy.IdValidationOpts
	next    validator
}

func (rv *requiredValidator) validate(f *field) {
	// required validation option is mandatory to start with the validation
	// if the required option is nil, return right away
	if rv.options.Required == nil {
		return
	}

	isRequired := *rv.options.Required
	// required == false --> we validate this field when non-empty, this field is not required
	if !isRequired && isEmpty(f.value) {
		return
	}

	// required == true ---> we validate this field, this field must own non-empty value
	if isRequired && isEmpty(f.value) {
		f.validationError = fmt.Errorf("'%s' cannot be empty", f.name)
		return
	}

	rv.next.validate(f)
}

func (rv *requiredValidator) setNext(v validator) {
	rv.next = v
}

func isEmpty(in interface{}) bool {
	switch val := in.(type) {
	case string:
		if len(strings.TrimSpace(val)) != 0 {
			return false
		}
	case *timestamppb.Timestamp:
		if val != nil {
			return false
		}
	}

	return true
}
