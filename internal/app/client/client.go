package client

import (
	"bufio"
	"fmt"
	"net"
)

func Client() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		println("Error")
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	println(status)
}
