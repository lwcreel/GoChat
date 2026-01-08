package server

import "net"

func Server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		println("Error")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			println("Error")
		}
		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {

	println("Connected To: ")
	println(connection.LocalAddr().String())

}
