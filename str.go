package str

import (
	"regexp"
)

func IsAlphaNum(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(str)
}
