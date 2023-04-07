package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"strings"
)

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unfading := int(plantText[length-1])
	return plantText[:(length - unfading)]
}

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	latest := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, latest...)
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
	encrypted := make([]byte, len(origData))
	blockModel.CryptBlocks(encrypted, origData)
	return strings.ToUpper(hex.EncodeToString(encrypted))
}
