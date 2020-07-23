package utils

import (
	"fmt"
	"strings"
)

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array),"[]"),"","",-1)
}

func ByteToString(b []byte) string {
	return string(b[:])
}

func StringToByte(s string) []byte {
	return []byte(s)
}
