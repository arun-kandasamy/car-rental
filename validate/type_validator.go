package validate

import (
	"fmt"
	"strings"

	policy "github.com/example/car-rental-service/proto/generated/policy"
)

type typeValidator struct {
	options *policy.IdValidationOpts
	next    validator
}

const (
	carIdPrefix  = "car_"
	gargIdPrefix = "gar_"
)

func (tv *typeValidator) validate(f *field) {
	// no type specified; we check only for required option in this case
	if tv.options.Type == nil {
		return
	}

	// only string values are supported
	v, ok := f.value.(string)
	if !ok {
		return
	}

	fieldType := *tv.options.Type
	switch fieldType {
	case "car":
		if !strings.HasPrefix(v, carIdPrefix) {
			f.validationError = fmt.Errorf("car id '%s' not valid", v)
			return
		}
	case "garage":
		if !strings.HasPrefix(v, gargIdPrefix) {
			f.validationError = fmt.Errorf("garage id '%s' not valid", v)
			return
		}
	default:
		f.validationError = fmt.Errorf("fieldtype '%s' not known", fieldType)
	}

	tv.next.validate(f)
}

func (dv *typeValidator) setNext(v validator) {
	dv.next = v
}
