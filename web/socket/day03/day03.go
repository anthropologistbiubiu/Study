package day03

import (
	"bufio"
	"fmt"
	"net"
)

func server() {
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go func(net.Conn) {
			for {
				input := bufio.NewReader(conn)
				recv, err := input.ReadString('\n')
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("from %+v content:%+v", conn.RemoteAddr(), recv)
				send := "ok"
				conn.Write([]byte(send))
			}
		}(conn)
	}
}
