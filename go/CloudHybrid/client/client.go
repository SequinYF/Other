package main

import (
	"net"
	"fmt"
	"log"
	"reflect"
	"io"
	"os"
	"encoding/json"
)

type FileInfo struct {
	Name string
	Size int
	Message []byte
}

const defaultBufSize = 4096
const server_addr = "127.0.0.1:9000"

func main()  {
	conn, err := net.Dial("tcp", server_addr)
	if err != nil {
		log.Fatal("cilent dail")
	}
	//
	fmt.Println("connect to server: " + conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr())
	fmt.Println(reflect.TypeOf(conn.LocalAddr()))
//	conn.Write([]byte("from client"))
	buf := make([]byte, 10)
	upload_to_server(conn, filepath)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			conn.Close()
		}

		fmt.Print(string(buf[:n]))
	}
}

func checkFileIsExist(filename string) bool {
	var exist bool = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

const filepath  =  "/Users/sequin_yf/go/src/CloudHybrid/fds"

func upload_to_server(conn net.Conn, filepath string) bool {

	var file *os.File

	err := checkFileIsExist(filepath)
	if err != true {
		log.Fatal("file not exist")
	}
	file, er := os.Open(filepath)
	if er != nil {
		log.Println(filepath + "cant open")
	}

	var temp_size int = 0
	temp_file := make([]byte, defaultBufSize)

	for {
		n, _ := file.Read(temp_file)
		if 0 == n {
			break
		}
		temp_size += n
	}

	fi := &FileInfo {
		filepath,
		temp_size,
		temp_file,
	}

	b, er := json.Marshal(fi)
	if er != nil {
		log.Fatal("marshal")
	}
	fmt.Println(string(b))
	conn.Write(b)
	return true
}