package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

var randReader io.Reader = rand.Reader

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func NewMD5() string {
	h := md5.New()
	return hex.EncodeToString(h.Sum(nil))
}

func RandSHA1() string {
	h := sha1.New()
	buf := make([]byte, 16)
	io.ReadFull(randReader, buf)
	h.Write(buf)

	return hex.EncodeToString(h.Sum(nil))
}
