package ipc

import (
	"testing"
)

type EchoServer struct {
	Server
}

func (server *EchoServer) Handle(request string) string {
	return "ECHO:" + request
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("addplayer", "From client1")
	resp2, _ := client2.Call("help", "From client2")

	if resp1.Body != "ECHO:From Client1" || resp2.Body != "ECHO:From Client2" {
		t.Error("IpcClient.Call faild. resp1", resp1, "resp2", resp2)
	}

	client1.Close()
	client2.Close()
}