package base64url

/*
	Implementation of base64url removing the trailing = and also changing +,/ to the more
	url friendly -,_ characters.

	The decode implementation is reversing the process before running DecodeString from the base64 package.
*/

import (
	"encoding/base64"
	"errors"
	"strings"
)

func Encode(data []byte) string {
	str := base64.StdEncoding.EncodeToString(data)
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "/", "_", -1)
	str = strings.Replace(str, "=", "", -1)
	return str
}

func Decode(str string) ([]byte, error) {
	if strings.ContainsAny(str, "+/") {
		return nil, errors.New("invalid base64url encoding")
	}
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)
	for len(str)%4 != 0 {
		str += "="
	}
	return base64.StdEncoding.DecodeString(str)
}
