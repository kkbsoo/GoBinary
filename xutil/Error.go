package xutil

import (
	"log"
	"io"
)
func ErrorNilCheck(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func ErrorEOFCheck(e error)int {
	if e == io.EOF {
		return -1
	}
	return 1
}