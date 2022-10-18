package main

import (
	"fmt"
	"net"
)

func main() {

	i := 0
	for i < 3 {
		i++

		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			// handle error
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		//status, err := bufio.NewReader(conn).ReadString('\n')
	}
}
