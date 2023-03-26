package validate

import (
	"fmt"

	policy "github.com/example/car-rental-service/proto/generated/policy"
)

type lengthValidator struct {
	options *policy.IdValidationOpts
	next    validator
}

func (lv *lengthValidator) validate(f *field) {
	if lv.options.Length == nil {
		return
	}

	lenV := *lv.options.Length
	v, ok := f.value.(string)
	if !ok {
		return
	}

	if len(v) != int(lenV) {
		f.validationError = fmt.Errorf("'%s' length is not equal to %d, but %d", f.name, lenV, len(v))
		return
	}

}

func (lv *lengthValidator) setNext(v validator) {
	lv.next = v
}
