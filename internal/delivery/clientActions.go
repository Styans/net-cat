package delivery

import (
	"bufio"
	"net"
	"sync"
)

var (
	Joining      chan Client
	Lefting      chan Client
	Messege      chan Client
	Listeners    map[string]*net.Conn
	Mut          sync.Mutex
	MaxListeners = 0
)

type Client struct {
	Messege string
	Name    string
}

func ProcessClient(conn net.Conn) {
	b := make([]byte, 1024)
	n, _ := conn.Read(b)
	c := Client{
		Name: string(b[:n-1]),
	}
	Listeners[string(b[:n-1])] = &conn
	Joining <- c

	defer conn.Close()
	sc := bufio.NewScanner(conn)

	for sc.Scan() {

		msg := sc.Text()

		c.Messege = msg
		Messege <- c
	}
}
