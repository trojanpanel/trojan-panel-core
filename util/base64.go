package util

import (
	"encoding/base64"
)

// Base64Decode2 解密两次
func Base64Decode2(s string) (string, error) {
	decodeUser, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	decodeString, err := base64.StdEncoding.DecodeString(string(decodeUser))
	if err != nil {
		return "", err
	}
	return string(decodeString), nil
}

// Base64Encode2 加密两次
func Base64Encode2(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(s))))
}
