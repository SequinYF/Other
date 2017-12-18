package main
import (
	"net"
	"fmt"
	"bufio"
	"time"
)

const ip = "127.0.0.1:9000"
var sig chan bool

func main() {
	addr,err := net.ResolveTCPAddr("tcp", ip)
	if err != nil {
		return
	}

	conn, err := net.DialTCP("tcp", nil,  addr)
	if err != nil {
		return
	}

	defer conn.Close()

	fmt.Println("Connected!")

	go recive(conn)

	b := []byte("message\n")
	conn.Write(b)
	<-sig
}

func recive(conn *net.TCPConn) {

	read := bufio.NewReader(conn)
	for{
		message, err := read.ReadString('\n')
		if err != nil {
			sig <- true
			break
		}
		fmt.Println("recive:" + message)
		time.Sleep(time.Second)
		mes :=  "message" + "\n"
		b := []byte(mes)
		conn.Write(b)
	}

}