package validators

import v "github.com/gima/govalid/v1"

func Contact() v.Validator {
	return v.Object(
		v.ObjKV("name", v.String(v.StrMin(1))),
		v.ObjKV("url", v.String(StringUrl())),
		v.ObjKV("email", v.String(StringEmail())),
	)
}

