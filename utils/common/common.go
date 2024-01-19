package common

import (
	"log"
	"strconv"
)

func ParseUint(str string) uint {
	resultStatus, err := strconv.ParseUint(str, 16, 0)
	if err != nil {
		log.Fatalln(err)
		return 0
	}
	return uint(resultStatus)
}
