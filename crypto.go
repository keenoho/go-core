package core

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rc4"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func DecryptAes(str string, key string) string {
	ciphertext, _ := hex.DecodeString(strings.ToUpper(str))
	pkey := []byte(key)
	block, _ := aes.NewCipher(pkey)
	blockModel := cipher.NewCBCDecrypter(block, pkey)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText)
	return string(plantText)
}

func EncryptAes(str string, key string) string {
	origData := []byte(str)
	origData = PKCS7Padding(origData)
	pkey := []byte(key)
	block, _ := aes.NewCipher(pkey)
	blockModel := cipher.NewCBCEncrypter(block, pkey)
	crypted := make([]byte, len(origData))
	blockModel.CryptBlocks(crypted, origData)
	return strings.ToUpper(hex.EncodeToString(crypted))
}

func EncryptHMACSHA1(str string, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(str))
	res := hex.EncodeToString(mac.Sum(nil))
	return res
}

func EncryptMd5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptRC4(str string, key string) string {
	dest := make([]byte, len(str))
	cipher, _ := rc4.NewCipher([]byte(key))
	cipher.XORKeyStream(dest, []byte(str))
	return strings.ToUpper(hex.EncodeToString(dest))
}

func DecryptRC4(str string, key string) string {
	ciphertext, _ := hex.DecodeString(strings.ToUpper(str))
	dest := make([]byte, len(ciphertext))
	cipher, _ := rc4.NewCipher([]byte(key))
	cipher.XORKeyStream(dest, []byte(ciphertext))
	return string(dest)
}
