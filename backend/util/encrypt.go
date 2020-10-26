package util

import (
	"encoding/hex"
	"crypto/sha1"
)


func Encrypt(password []byte) string{
	h := sha1.New()
	h.Write(append(password))
	vector:=hex.EncodeToString(h.Sum(nil))
	//vector:=h.Sum([]byte{})
	return string(vector)
}