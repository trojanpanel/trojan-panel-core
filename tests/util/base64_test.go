package util

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	auth := base64.StdEncoding.EncodeToString([]byte("123123"))
	fmt.Println(auth)
}
