package util

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

const saltSize = 16

func Sha1String(plain string) string {
	buf := make([]byte, saltSize, saltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		fmt.Println("random read failed ->", err)
	}

	h := sha1.New()
	h.Write(buf)
	h.Write([]byte(plain))

	return base64.URLEncoding.EncodeToString(h.Sum(buf))
}

// Sha1Match match
func Sha1Match(secret, plain string) bool {
	data, _ := base64.URLEncoding.DecodeString(secret)
	if len(data) != saltSize+sha1.Size {
		fmt.Println("wrong length of data")
		return false
	}
	h := sha1.New()
	h.Write(data[:saltSize])
	h.Write([]byte(plain))
	return bytes.Equal(h.Sum(nil), data[saltSize:])
}

func SHA224String(password string) string {
	hash := sha256.New224()
	hash.Write([]byte(password))
	val := hash.Sum(nil)
	str := ""
	for _, v := range val {
		str += fmt.Sprintf("%02x", v)
	}
	return str
}
