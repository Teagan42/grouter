package validators

import (
	v "github.com/gima/govalid/v1"
)

func SwaggerDoc() v.Validator {
	scheme := func () v.Validator {
		return v.Or(
			v.String(v.StrIs("http")),
			v.String(v.StrIs("https")),
			v.String(v.StrIs("ws")),
			v.String(v.StrIs("wss")),
		)
	}

	return v.Object(
		v.ObjKV("swagger", v.String(v.StrIs("2.0"))),
		v.ObjKV("info", Info()),
		v.ObjKV("host", v.Optional(v.String())), // TODO
		v.ObjKV("basePath", v.Optional(v.String())), // TODO
		v.ObjKV("schemes", v.Optional(
				v.Array(v.ArrEach(scheme())),
		)),
		v.ObjKV("consumes", v.Optional(
			v.Array(v.ArrEach(v.String(StringMimeType()))),
		)),
		v.ObjKV("produces", v.Optional(
			v.Array(v.ArrEach(v.String(StringMimeType()))),
		)),
		v.ObjKV("paths", v.Optional(
			Paths(),
		)),
		v.ObjKV("securityDefinitions", v.Optional(
			SecurityDefinition(),
		)),
	)
}

