package server

import (
	"net-cat/internal/delivery"
	"time"
)

const (
	join = " has joined our chat...\n"
	left = " has left our chat...\n"
)

func BroadCastServer() {
	var msg delivery.Client
	for {
		select {
		case msg = <-delivery.Joining:
			msg.Messege = "\n" + msg.Name + join
		case <-delivery.Lefting:
			msg.Messege = "\n" + msg.Name + left
		case msg = <-delivery.Messege:
			msg.Messege = "\n[" + msg.Time + "][" + msg.Name + "]:" + msg.Messege + "\n"

		}
		anounce(msg)
	}
}

func anounce(msg delivery.Client) {
	delivery.Mut.Lock()
	for name, el := range delivery.Listeners {
		if name != msg.Name {
			(*el).Write([]byte(msg.Messege))
			(*el).Write([]byte("[" + time.Now().Format(delivery.TimeFormat) + "][" + name + "]:"))
		}
	}
	delivery.Mut.Unlock()
}
