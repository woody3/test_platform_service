package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func StringJoin(args ...string) string {
	var builder strings.Builder
	for _, s := range args {
		builder.WriteString(s)
	}
	return builder.String()
}

func EncryptMD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
