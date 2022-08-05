package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/module/constant"
)

var aesKey = []byte(constant.SaltKey)

// AesEncode 加密
func AesEncode(origData string) (string, error) {
	pass := []byte(origData)
	xPass, err := aesEncrypt(pass, aesKey)
	if err != nil {
		logrus.Errorf("aes 加密错误 err: %v\n", err)
		return "", errors.New(constant.SysError)
	}
	return base64.StdEncoding.EncodeToString(xPass), nil
}

// AesDecode 解密
func AesDecode(crypted string) (string, error) {
	bytesPass, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		logrus.Errorf("base64 解密错误 err: %v\n", err)
		return "", errors.New(constant.SysError)
	}
	tPass, err := aesDecrypt(bytesPass, aesKey)
	if err != nil {
		logrus.Errorf("aes 解密错误 err: %v\n", err)
		return "", errors.New(constant.SysError)
	}
	return string(tPass), nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func aesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func aesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS5UnPadding(origData)
	return origData, nil
}
