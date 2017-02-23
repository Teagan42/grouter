package validators


import (
	v "github.com/gima/govalid/v1"
)

func HttpMethod() v.Validator {
	return v.Or(
		v.String(v.StrIs("put")),
		v.String(v.StrIs("post")),
		v.String(v.StrIs("get")),
		v.String(v.StrIs("delete")),
	)
}

func Path() v.Validator {
	return v.Object(
		v.ObjKeys(HttpMethod()),
		v.ObjValues(Operation()),
	)
}
