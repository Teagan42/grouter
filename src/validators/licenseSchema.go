package validators

import v "github.com/gima/govalid/v1"

func License() v.Validator {
	return v.Object(
		v.ObjKV("name", v.String(v.StrMin(1))),
		v.ObjKV("url", v.Optional(v.String(StringUrl()))),
	)
}