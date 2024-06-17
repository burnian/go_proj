package utils

import (
	"strings"
)

func StrSplit(str string, sep string) []string {
	return strings.Split(str, sep)
}
