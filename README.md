# Build TCP Chat with GO (Golang) 🚀

running with your netcat localhost on `port`: 
```
go run cmd/main.go
```
connect to localhost (netcat):
```
nc localhost 9090
```
connect to localhost (telnet): 
```
telnet localhost 9090
```

command line for your `client`:
- `/nick <name>` - get a name, otherwise user will stay anonymous.
- `/join <name>` - join a room, if room doesn't exist, the new room will be created. User can be only in one room at the same time.
- `/rooms`       - show list of available rooms to join.
- `/msg	<msg>` - roadcast message to everyone in a room.
- `/quit` - disconnects from the chat server.
