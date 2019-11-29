package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

// found in https://github.com/bediger4000/udp-echo-server/blob/master/server.go

func main() {
	portS, exists := os.LookupEnv("PORT")
	var port int = 30001 // default
	var err error
	if exists {
		port, err = strconv.Atoi(portS)
		if err != nil {
			log.Fatal("Cannot read ENV port")
		}
	}

	addr := net.UDPAddr{Port: port, IP: net.ParseIP("0.0.0.0")}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 2048)

	for {
		fmt.Printf("Accepting a new packet\n")

		cc, remote, rderr := conn.ReadFromUDP(b)

		if rderr != nil {
			fmt.Printf("net.ReadFromUDP() error: %s\n", rderr)
		} else {
			fmt.Printf("Read %d bytes from socket\n", cc)
			fmt.Printf("Bytes: %q\n", string(b[:cc]))
		}

		fmt.Printf("Remote address: %v\n", remote)

		cc, wrerr := conn.WriteTo(b[0:cc], remote)
		if wrerr != nil {
			fmt.Printf("net.WriteTo() error: %s\n", wrerr)
		} else {
			fmt.Printf("Wrote %d bytes to socket\n", cc)
		}
	}

	fmt.Printf("Out of infinite loop\n")
}
