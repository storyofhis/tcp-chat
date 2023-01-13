package controllers

type CommandID int

const (
	CMD_NICK CommandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type Command struct {
	Id     CommandID
	Client *Clients
	Args   []string
}
