package convert

import (
	"strconv"
	"strings"
)

func MakeInt(s string) int {
	s = strings.TrimSpace(s)
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
