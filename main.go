package main

import (
	"lib/lib"
	"lib/server"
)

func main() {
	library := lib.NewLibrary()
	handlers := server.NewHandlers(library)
	server := server.NewServer(handlers)

	server.StartServer()
}
