package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptMd5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}
