package main

import (
	"fmt"
	"github.com/volf52/dnsgo/pkg/dns"
	"io/ioutil"
	"net"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	queryBytes, err := ioutil.ReadFile("./test_data/query_packet")
	HandleErr(err)

	buff := dns.BufferFrom(queryBytes)
	header := dns.ParseHeaderFrom(buff)
	lbl := dns.ParseLabelSequenceFrom(buff)

	fmt.Println(header)
	fmt.Println(lbl)

	fmt.Printf("Read %d bytes\n", len(queryBytes))

	dnsAddr, err := net.ResolveUDPAddr("udp4", "8.8.8.8:53")
	HandleErr(err)

	sock, err := net.DialUDP("udp", nil, dnsAddr)
	HandleErr(err)
	defer CloseSocket(sock)
	fmt.Println("Connected!")

	sent, err := sock.Write(queryBytes)
	HandleErr(err)

	fmt.Printf("Sent %d bytes...\n", sent)

	buffer := make([]byte, 1024)
	n, err := sock.Read(buffer)
	HandleErr(err)

	fmt.Printf("Received %d bytes...\n", n)

	err = ioutil.WriteFile("lala", buffer[:n], 0644)
	HandleErr(err)
}

func CloseSocket(sock *net.UDPConn) {
	func(conn *net.UDPConn) {
		err := conn.Close()
		HandleErr(err)
	}(sock)
}
