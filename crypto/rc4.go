package crypto

import (
	"crypto/rc4"
	"encoding/hex"
	"strings"
)

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
