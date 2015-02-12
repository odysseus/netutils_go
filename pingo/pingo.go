package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	flag.Parse()

	straddr := flag.Arg(0)
	pid := os.Getpid()

	raddr, err := net.ResolveIPAddr("ip4:icmp", straddr)
	checkErr(err)

	conn, err := net.DialIP("ip4:icmp", nil, raddr)
	checkErr(err)

	fmt.Println(conn)
	fmt.Println(pid)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func makePacket(id, seq int, payload []byte) []byte {
	paylen := len(payload)
	p := make([]byte, 8+paylen)

	p[0] = 8 // Type
	p[1] = 0 // Code
	p[2] = 0 // Header checksum
	p[3] = 0 // Header checksum
	p[4] = 0 // Identifier
	p[5] = 0 // Identifier
	p[6] = 0 // Sequence number
	p[7] = 0 // Sequence number

	for i := 0; i < paylen; i++ {
		p[8+i] = payload[i]
	}

	return p
}
