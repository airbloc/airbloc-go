package testutils

import (
	"net"
	"time"
)

// ReservePort automatically reserves available port in system.
func ReservePort() int {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	if err := lis.Close(); err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)
	return lis.Addr().(*net.TCPAddr).Port
}
