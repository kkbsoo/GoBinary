package xutil

import (
	"net"
)

func OpenUDP(sPort string) (*net.UDPConn, error){

	ServerAddr,err := net.ResolveUDPAddr("udp",sPort)
	ErrorNilCheck(err)
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	ErrorNilCheck(err)
	return ServerConn, err
}

func RecvUDP(sBuf string) int{
	return 1
}
