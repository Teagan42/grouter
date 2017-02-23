package validators

import (
	v "github.com/gima/govalid/v1"
	"fmt"
)

// http://regexlib.com/REDetails.aspx?regexp_id=3755

var urlRelativePattern = "\\/((?:[a-zA-Z\\d_\\.\\/])|(?:%[a-fA-F0-9]{2,}))*\\/?"
var urlPattern = fmt.Sprintf("((http(s)?\\:\\/\\/)?([a-zA-Z0-9]{1}[a-zA-Z0-9\\.\\-_]{0,255})(\\:\\d{1,5})?){1}(%s)?", urlRelativePattern)

func StringRelative() v.StringOpt {
	return v.StrRegExp(fmt.Sprintf("^%s$", urlRelativePattern))
}

func StringUrl() v.StringOpt {
	return v.StrRegExp(fmt.Sprintf("^%s$", urlPattern))
}
