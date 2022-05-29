package main

import (
	"TZapigendoc/internal/server"
)

func main() {
	serv := server.NewApp()
	serv.Run()
}
