package server

import (
	"net-cat/internal/delivery"
)

func BroadCastServer() {
	for {
		select {
		case msg := <-delivery.Joining:
			msg.Messege = msg.Name + " has joined our chat...\n"
			anounce(msg)
		case <-delivery.Lefting:
		case msg := <-delivery.Messege:
			msg.Messege = msg.Messege + "\n"
			anounce(msg)
		}
	}
}

func anounce(msg delivery.Client) {
	delivery.Mut.Lock()
	for name, el := range delivery.Listeners {
		if name != msg.Name {
			(*el).Write([]byte(msg.Messege))
		}
	}
	delivery.Mut.Unlock()
}
