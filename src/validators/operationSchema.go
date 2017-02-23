package validators

import (
	v "github.com/gima/govalid/v1"
)

func StringMimeType() v.StringOpt {
	const mimeTypePattern = "^(application|audio|example|image|message|model|multipart|text|video)\\/[a-zA-Z0-9\\*]+([+.-][a-zA-z0-9]+)*$"

	return v.StrRegExp(mimeTypePattern)
}

func Operation() v.Validator {
	return v.Object(
		v.ObjKV("tags",
			v.Optional(
				v.Array(v.ArrEach(
					v.String(v.StrMin(1)))))),
		v.ObjKV("summary",
			v.String(
				v.StrMin(1),
				v.StrMax(120))),
		v.ObjKV("description", v.String(v.StrMin(1))),
		v.ObjKV("consumes", v.Optional(v.Array(v.ArrEach(v.String(StringMimeType()))))), // TODO
		v.ObjKV("produces", v.Optional(v.Array(v.ArrEach(v.String(StringMimeType()))))),
		v.ObjKV("parameters", v.Optional(v.Array(v.ArrEach(Parameter())))),
		v.ObjKV("schemes", v.Optional(v.Array(v.ArrEach(v.Or(
			v.String(v.StrIs("http")),
			v.String(v.StrIs("https")),
		))))),
		v.ObjKV("deprecated", v.Boolean()),
	)
}
