package util

import "encoding/base64"

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
