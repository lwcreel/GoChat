package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lwcreel/GoChat/internal/app/client"
	"github.com/lwcreel/GoChat/internal/app/server"
)

func main() {
	if os.Args[1] == "server" {
		fmt.Println("Starting Server...")
		http.HandleFunc("/", server.Server)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	} else {
		fmt.Println("Starting Client...")
		client.Client()
	}
}
