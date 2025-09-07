package string

import (
	"strings"
)



func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsEqual(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

