package controllers

import "net"

type Room struct {
	Name    string
	Members map[net.Addr]*Clients
}

func (r *Room) broadcast(sender *Clients, msg string) {
	for addr, m := range r.Members {
		if sender.Conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}
