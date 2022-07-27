package utils

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/martinlindhe/base36"
	"strings"
)

var (
	replacer = strings.NewReplacer(
		" ", "0",
		"O", "0",
		"I", "1",
	)
)

func Err2Hash(err error) string {
	u64 := hash(err.Error())
	return encode(u64)
}

func Err2Hashcode(err error) (uint64, string) {
	u64 := hash(err.Error())
	codeStr := encode(u64)
	u64, _ = decode(codeStr)
	return u64, codeStr
}

func encode(code uint64) string {
	s := fmt.Sprintf("%4s", base36.Encode(code))
	return replacer.Replace(s)
}

func decode(s string) (uint64, bool) {
	if len(s) != 4 {
		return 0, false
	}
	s = strings.Replace(s, "l", "1", -1)
	s = strings.ToUpper(s)
	s = replacer.Replace(s)
	code := base36.Decode(s)
	return code, code > 0
}

// hash 函数可以自定义
func hash(s string) uint64 {
	h := md5.Sum([]byte(s))
	u := binary.BigEndian.Uint32(h[0:16])
	return uint64(u & 0xFFFFF)
}
