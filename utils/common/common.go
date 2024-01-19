package common

import (
	"log"
	"strconv"
	"strings"
)

func ParseUint(str string) uint {
	s := strings.TrimPrefix(str, "0x")
	resultStatus, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		log.Fatalln(err)
		return 0
	}
	return uint(resultStatus)
}
