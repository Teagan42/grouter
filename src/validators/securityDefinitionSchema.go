package validators

import (
	v "github.com/gima/govalid/v1"
)

func SecurityDefinition() v.Validator {
	return v.Object(
		v.ObjKeys(
			v.String(
				v.StrRegExp("^[a-zA-Z0-9_]*$"))),
		v.ObjValues(
			v.Object(
				v.ObjKV("type", v.Or(
					v.String(v.StrIs("basic")),
					v.String(v.StrIs("apiKey")),
				)),
				v.ObjKV("description", v.Optional(v.String())),
				v.ObjKV("name", v.String(v.StrMin(1))),
				v.ObjKV("in", v.Or(
					v.String(v.StrIs("header")),
					v.String(v.StrIs("query")),
				)),
				// TODO : Oauth2 Keys, conditionally required
			),
		),
	)
}

