package main

import (
	"fmt"
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp4", ":8829")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	fmt.Println("Listening for UDP messages on port 8829...")

	buf := make([]byte, 1024)

	for {
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %s: %s\n", addr, buf[:n])
	}
}
