package server

import (
	"errors"
	"fmt"
	"net"
	"net-cat/internal/delivery"
	"net-cat/internal/helpers"
	"net-cat/internal/myerrors"
	"os"
)

var (
	port    = "localhost:8080"
	errLogo = "( ͡° ͜ʖ ͡°)\n"
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
		return errors.New(myerrors.IncorectArgs)
	}
	delivery.Joining = make(chan delivery.Client)
	delivery.Lefting = make(chan delivery.Client)
	delivery.Messege = make(chan delivery.Client)
	delivery.Listeners = make(map[string]*net.Conn)
	return nil
}

func StartServer(maxListeners int) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Printf("Listening for connections on %s\n", listener.Addr().String())
	go BroadCastServer()
	logo, err := os.ReadFile("logo.txt")
	if err != nil {
		logo = []byte(errLogo)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		delivery.Mut.Lock()
		if delivery.NListeners > maxListeners {
			conn.Write([]byte(myerrors.MaxConnectionsMessage))
			conn.Close()
			delivery.Mut.Unlock()

			continue
		}
		delivery.NListeners++
		delivery.Mut.Unlock()

		go delivery.ProcessClient(logo, conn)
	}
}
