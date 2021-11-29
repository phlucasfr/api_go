package main

import "github.com/phlucasfr/api_go/server"

func main() {
	server := server.NewServer()

	server.Run()
}
