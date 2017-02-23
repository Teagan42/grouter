package validators

import v "github.com/gima/govalid/v1"

func Info() v.Validator{
	return v.Object(
		v.ObjKV("title", v.String(v.StrMax(120), v.StrMin(1))),
		v.ObjKV("description", v.String()),
		v.ObjKV("termsOfService", v.Optional(v.String())),
		v.ObjKV("contact", v.Optional(Contact())),
		v.ObjKV("license", v.Optional(License())),
	)
}
