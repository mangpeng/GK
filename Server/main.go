package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		fmt.Println("Wait for clients")
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Connected from a new client!!")
	fmt.Println(conn.LocalAddr().String())
	fmt.Println(conn.LocalAddr().Network())
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.RemoteAddr().Network())
}
