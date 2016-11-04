package xutil

import (
	"fmt"
)

func GetIpStr(iIp[4] int32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(iIp[0]), byte(iIp[0]>>8), byte(iIp[0]>>16), byte(iIp[0]>>24))
}