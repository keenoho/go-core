package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func EncryptHMACSHA1(str string, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(str))
	res := hex.EncodeToString(mac.Sum(nil))
	return res
}
