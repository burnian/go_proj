package utils

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("utils.go init()")
}

func StrSplit(str string, sep string) []string {
	return strings.Split(str, sep)
}
