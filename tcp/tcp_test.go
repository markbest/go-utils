package utils

import (
	"net"
	"testing"
)

var (
	addr     = "127.0.0.1"
	port     = "1234"
	maxRetry = 10
)

func TestTcp(t *testing.T) {
	server := NewTCPServer(addr, port, maxRetry)
	lis, lisErr := server.Listen()
	if lisErr != nil {
		t.Error(lisErr.Error())
	}

	// server listen
	readerChannel := make(chan []byte, 16)
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				continue
			}
			go handleConnection(conn, readerChannel)
		}
	}()

	// client sender msg
	client := NewTCPClient(addr, port, maxRetry)
	if handleErr := client.ReadWrite(sendMsg); handleErr != nil {
		t.Error(handleErr)
	}
	client.Close()

	rs := string(<-readerChannel)
	t.Log(rs)
}

func handleConnection(conn net.Conn, readerChannel chan []byte) {
	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			return
		}
		readerChannel <- buffer
	}
}

func sendMsg(conn *net.TCPConn) error {
	conn.Write([]byte("Hello World!"))
	return nil
}
