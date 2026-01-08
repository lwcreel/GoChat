package main

import (
	"fmt"
	"os"

	"github.com/lwcreel/GoChat/internal/app/client"
	"github.com/lwcreel/GoChat/internal/app/server"
)

func main() {
	if os.Args[1] == "server" {
		fmt.Println("Starting Server...")
		server.Server()
	} else {
		fmt.Println("Starting Client...")
		client.Client()
	}
}
