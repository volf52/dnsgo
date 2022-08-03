package main

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/volf52/dnsgo/pkg/dns/packet"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	queryPacket := packet.IpQuery("api.carbonteq-livestream.ml")
	queryBytes := queryPacket.Bytes()

	dnsAddr, err := net.ResolveUDPAddr("udp4", "8.8.8.8:53")
	HandleErr(err)

	sock, err := net.DialUDP("udp", nil, dnsAddr)
	HandleErr(err)
	defer CloseSocket(sock)
	fmt.Println("Connected!")

	sent, err := sock.Write(queryBytes)
	HandleErr(err)

	fmt.Printf("Sent %d bytes...\n", sent)

	buff := make([]byte, 1024)
	n, err := sock.Read(buff)
	HandleErr(err)

	fmt.Printf("Received %d bytes...\n", n)

	resp := packet.FromResponse(buff[:n])

	fmt.Print(resp)

	err = ioutil.WriteFile("lala", buff[:n], 0644)
	HandleErr(err)
}

func CloseSocket(sock *net.UDPConn) {
	err := sock.Close()
	HandleErr(err)
}
