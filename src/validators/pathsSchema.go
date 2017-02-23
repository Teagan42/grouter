package validators

import (
	v "github.com/gima/govalid/v1"
)

func Paths() v.Validator {
	return v.Object(
		v.ObjKeys(
			v.String(
				v.StrRegExp("\\/(?:(?:[a-zA-Z0-9_]+)|(?:\\:[a-zA-Z0-9_]+)\\/?)?"))),
		v.ObjValues(Path()),
	)
}
