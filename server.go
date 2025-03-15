package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	pc, err := net.ListenPacket("udp4", ":8830")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	// Enable SO_BROADCAST
	rawConn, err := pc.(*net.UDPConn).SyscallConn()
	if err != nil {
		panic(err)
	}
	rawConn.Control(func(fd uintptr) {
		syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1)
	})

	addr, err := net.ResolveUDPAddr("udp4", "192.168.1.255:8829")
	if err != nil {
		panic(err)
	}

	_, err = pc.WriteTo([]byte("data to transmit"), addr)
	if err != nil {
		panic(err)
	}

	fmt.Print("Sent stuff...")
}
