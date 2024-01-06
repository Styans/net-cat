package delivery

import (
	"bufio"
	"net"
	"net-cat/internal/helpers"
	"net-cat/internal/myerrors"
	"strings"
	"sync"
	"time"
)

var (
	Joining    chan Client
	Lefting    chan Client
	Messege    chan Client
	Listeners  map[string]*net.Conn
	Mut        sync.Mutex
	NListeners = 0
)

const (
	meeting    = "Welcome to TCP-Chat!\n"
	signIn     = "[ENTER YOUR NAME]: "
	TimeFormat = "2006-01-02 15:04:05"
)

type Client struct {
	Messege string
	Time    string
	Name    string
}

func ProcessClient(logo []byte, conn net.Conn) {
	conn.Write([]byte(meeting))
	conn.Write(logo)

	name := getClientName(conn)
	c := Client{
		Name: strings.Title(name),
	}
	Mut.Lock()
	Listeners[name] = &conn
	Joining <- c
	Mut.Unlock()

	defer conn.Close()
	sc := bufio.NewScanner(conn)
	conn.Write([]byte("[" + time.Now().Format(TimeFormat) + "][" + name + "]:"))

	for sc.Scan() {

		msg := sc.Text()
		c.Time = time.Now().Format(TimeFormat)
		c.Messege = msg
		Mut.Lock()

		Messege <- c
		Mut.Unlock()
		conn.Write([]byte("[" + time.Now().Format(TimeFormat) + "][" + name + "]:"))

	}
	closeConn()
	Mut.Lock()
	Lefting <- c
	delete(Listeners, name)
	Mut.Unlock()
}
func closeConn() {
	Mut.Lock()
	NListeners--
	Mut.Unlock()
}

func getClientName(conn net.Conn) string {
	b := make([]byte, 1024)
	var firstName string

L:
	for {
		conn.Write([]byte(signIn))
		n, _ := conn.Read(b)
		firstName = strings.Title(string(b[:n-1]))
		switch {
		case !helpers.IsAscii(firstName):
			conn.Write([]byte(myerrors.NotAscii))

		case !cloneName(firstName, conn):
			conn.Write([]byte(myerrors.NickIsBusy))
		default:
			break L

		}

	}

	return firstName
}

func cloneName(firstName string, conn net.Conn) bool {
	for secondName, _ := range Listeners {
		if firstName == secondName {
			return false
		}
	}
	return true
}
