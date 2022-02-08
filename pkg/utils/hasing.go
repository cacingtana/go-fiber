package utils

import (
	"crypto/md5"
	"encoding/hex"
)

type Hash struct {
}

func (h *Hash) MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
