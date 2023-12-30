package str

import (
	"regexp"
)

func IsAlphaNum(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(str)
}

func TrimSides(str string, leftTrimCount int, rightTrimCount int) string {
	return (str[leftTrimCount:])[:(len(str) - (leftTrimCount + rightTrimCount))]
}
