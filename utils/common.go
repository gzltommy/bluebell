package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "zorro.tommy.miao"

func EncryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}
