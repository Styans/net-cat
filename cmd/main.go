package main

import (
	"fmt"
	"net-cat/internal/server"
)

func main() {

	err := server.InitVars()
	if err != nil {
		fmt.Println(err)
		return
	}
	server.StartServer(10)
}
