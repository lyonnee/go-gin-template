package util

import (
	"regexp"
)

func CompressJson(jsonStr string) string {
	regexpSource := `\n|\t|(\s)(\s)|(\s)(^[A-Z]|[a-z]|[0-9]|#|\+|-|\*\/|=)`

	result := regexp.MustCompile(regexpSource).ReplaceAllString(jsonStr, "")

	return result
}
