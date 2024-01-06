package server

import (
	"fmt"
	"net"
	"net-cat/internal/delivery"
	"net-cat/internal/helpers"
	"os"
)

var (
	port         = "localhost:8080"
	MaxListeners = 0
)

func InitVars() error {
	switch len(os.Args) {
	case 2:
		var err error
		port, err = helpers.CheckPort(os.Args[1])
		if err != nil {
			return err
		}
	case 1:
	default:

	}
	delivery.Joining = make(chan delivery.Client)
	delivery.Lefting = make(chan delivery.Client)
	delivery.Messege = make(chan delivery.Client)
	delivery.Listeners = make(map[string]*net.Conn)
	return nil
}

func StartServer() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	fmt.Printf("Listening for connections on %s\n", listener.Addr().String())
	go BroadCastServer()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		go delivery.ProcessClient(conn)
	}
}
