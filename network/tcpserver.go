package network

import (
	"fmt"
	"net"
)

func TcpServer(addr string) {
	tcp, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}

	listen, err := net.ListenTCP("tcp", tcp)

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("new conn", conn.RemoteAddr())

		go handler(conn)
	}

}

func handler(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 50)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("data:", string(buf))
}
