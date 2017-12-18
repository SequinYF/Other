package main

import (
	"net"
	"fmt"
	"bufio"
	//"time"
	"time"
)

const ipserver = "127.0.0.1:9000"

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ipserver)
	checkError(err)

	conn, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	defer conn.Close()

    fmt.Println("server start")
	for{
		client, err := conn.AcceptTCP()
		checkError(err)
		fmt.Println("client" + client.RemoteAddr().String())
		go echo(client)
	}
}

func echo(client *net.TCPConn) {

	defer client.Close()

	read := bufio.NewReader(client)

	for {
		message, err := read.ReadString('\n')
		checkError(err)
		fmt.Println("recive:" + message)
		tim := time.Now().String() + ": " + message
		fmt.Println("send:" + tim)
		b := []byte(tim)
		client.Write(b)
	}

}

func checkError(err error) {
	if err != nil {
		return
	}
}
