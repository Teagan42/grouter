package validators


import (
	v "github.com/gima/govalid/v1"
	"fmt"
	"reflect"
)

func ParameterType() v.Validator {
	var validatorValidator = func (data interface{}) (path string, err error) {
		var _, ok = data.(v.Validator)

		if !ok {
			return "validatorValidator", fmt.Errorf("expected validator, got %s", reflect.TypeOf(data))
		}

		return "", nil
	}
	return v.Function(validatorValidator)
}

func Parameter() v.Validator {
	return v.Object(
		v.ObjKV("name", v.String(v.StrMin(1))),
		v.ObjKV("in", v.Or(
			v.String(v.StrIs("query")),
			v.String(v.StrIs("header")),
			v.String(v.StrIs("path")),
			v.String(v.StrIs("formData")),
			v.String(v.StrIs("body")))),
		v.ObjKV("type", v.Object(
			v.ObjKV("name", v.String(v.StrMin(1))),
			v.ObjKV("validator", ParameterType()),
		)),
		v.ObjKV("description", v.String(v.StrMin(1))),
	)
}
