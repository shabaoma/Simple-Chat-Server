package main

import (
	"Simple-Chat-Application/message"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
)

type Connection struct {
	Id      string
	Conn    net.Conn
	ConnMap *map[string]net.Conn
}

func (c Connection) Process() {
	defer c.Conn.Close()

	for {
		buf := make([]byte, 1024)

		// 如果客户端不发送请求，携程就会一直卡在这里
		fmt.Printf("Wait for client: %v\n", c.Conn.RemoteAddr().String())
		n, err := c.Conn.Read(buf) // n: 发送的字节数
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
			} else {
				fmt.Println("Read message error:", err)
			}
			break
		}

		m := message.Message{}
		json.Unmarshal(buf[:n], &m)
		m.From = c.Id
		fmt.Printf("Read message %v\n", m)
		b, err := json.Marshal(m)
		if err != nil {
			fmt.Println("Marshal error:", err)
			continue
		}

		redirect_conn := (*c.ConnMap)[m.To]
		redirect_conn.Write(b)
	}
}

func main() {
	server, err := net.Listen("tcp", ":8080")
	connection_map := make(map[string]net.Conn)
	id := 0

	fmt.Println("Start server")
	if err != nil {
		fmt.Println("Start server error: ", err)
		return
	}
	defer server.Close()

	for {
		fmt.Println("Wait for client to connect")
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Accept error: ", err)
			continue
		}

		fmt.Printf("Assign id %v to %v\n", id, conn.RemoteAddr().String())
		str_id := strconv.Itoa(id)
		connection_map[str_id] = conn
		connection := Connection{str_id, conn, &connection_map}
		id++
		go connection.Process()
	}
}
