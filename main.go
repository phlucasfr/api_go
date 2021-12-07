package main

import "api_go/server"

func main() {
	server := server.NewServer()

	server.Run()
}
