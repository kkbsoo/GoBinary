package xutil

import (
	"bytes"
)

func Trim(sTrimBuf[] byte) string {
	return string(bytes.Trim(sTrimBuf[:], "\x00"))
}
