package controllers

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	Room     map[string]*Room
	Commands chan Command
}

func NewServer() *Server {
	return &Server{
		Room:     make(map[string]*Room),
		Commands: make(chan Command),
	}
}

func (server *Server) NewClient(Conn net.Conn) *Clients {
	log.Printf("new client has joined: %s\n", Conn.RemoteAddr().String())
	return &Clients{
		Conn:     Conn,
		Nick:     "Anonymous",
		Commands: server.Commands,
	}
}

func (server *Server) Run() {
	for cmd := range server.Commands {
		switch cmd.Id {
		case CMD_NICK:
			server.Nick(cmd.Client, cmd.Args)
		case CMD_JOIN:
			server.Join(cmd.Client, cmd.Args)
		case CMD_ROOMS:
			server.ListRooms(cmd.Client)
		case CMD_MSG:
			server.Msg(cmd.Client, cmd.Args)
		case CMD_QUIT:
			server.Quit(cmd.Client)
		}
	}
}

func (server *Server) Nick(c *Clients, args []string) {
	if len(args) < 2 {
		c.msg("nick is required. usage: /nick NAME")
		return
	}
	c.Nick = args[1]
	c.msg(fmt.Sprintf("all right, I will call you %s", c.Nick))
}

func (server *Server) Join(c *Clients, args []string) {
	if len(args) < 2 {
		c.msg("room name is required. usage: /join ROOM_NAME")
		return
	}

	roomName := args[1]

	r, ok := server.Room[roomName]
	if !ok {
		r = &Room{
			Name:    roomName,
			Members: make(map[net.Addr]*Clients),
		}
		server.Room[roomName] = r
	}

	r.Members[c.Conn.RemoteAddr()] = c
	server.QuitCurrentRoom(c)
	c.Room = r

	r.broadcast(c, fmt.Sprintf("%s joined the room", c.Nick))

	c.msg(fmt.Sprintf("welcome to %s", roomName))
}

func (s *Server) ListRooms(c *Clients) {
	var rooms []string
	for name := range s.Room {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("available rooms: %s", strings.Join(rooms, ", ")))
}

func (s *Server) Msg(c *Clients, args []string) {
	if len(args) < 2 {
		c.msg("message is required, usage: /msg MSG")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.Room.broadcast(c, c.Nick+": "+msg)
}

func (s *Server) Quit(c *Clients) {
	log.Printf("client has left the chat: %s", c.Conn.RemoteAddr().String())

	s.QuitCurrentRoom(c)

	c.msg("sad to see you go =(")
	c.Conn.Close()
}

func (s *Server) QuitCurrentRoom(c *Clients) {
	if c.Room != nil {
		oldRoom := s.Room[c.Room.Name]
		delete(s.Room[c.Room.Name].Members, c.Conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the room", c.Nick))
	}
}
