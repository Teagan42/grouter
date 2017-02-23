package validators

import (
	v "github.com/gima/govalid/v1"
)

var emailPattern = "^[a-zA-Z0-9\\._%+-]+@([[a-zA-Z0-9\\.-]+\\.[a-zA-Z]{2,})|(\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3})$"

func StringEmail() v.StringOpt {
	return v.StrRegExp(emailPattern)
}