package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	UserCh chan string
	Conn net.Conn
	Server *Server
} 

// factory for a new User instance
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		UserCh: make(chan string), // unbuffered channel
		Conn: conn,
		Server: server,
	}
	go user.PushMessageToClient() // start a goroutine to push message to the client
	
	return user
}

// get the message from the channel and send it to the client
func (u *User) PushMessageToClient() {
	for msg := range u.UserCh {
		u.Conn.Write([]byte(msg + "\n"))
	}
}

// send the message to the channel of the user,
// and the message will be sent to the channel of the server by PushMessageToClient() 
func (u *User) SendMessage(str string) {
	u.UserCh <- str // send the message to the channel
}

// user come online, add the user to the online user map of the server,
// and broadcast the user online status to all online users of the server.
func (u *User) Online() {
		// add user to the online user map of the server
		u.Server.mapLock.Lock()
		u.Server.OnlineMap[u.Name] = u
		u.Server.mapLock.Unlock()
		// broadcast the user online status to all online users of the server
		u.Server.BroadCast(u, " is online now. Welcome!")
}

// user go offline, remove the user from the online user map of the server,
// and broadcast the user offline status to all online users of the server.
func (u *User) Offline() {
	// remove user from the online user map of the server
	u.Server.mapLock.Lock()
	delete(u.Server.OnlineMap,u.Name)
	u.Server.mapLock.Unlock()
	// broadcast the user online status to all online users of the server
	u.Server.BroadCast(u, " is offline. Bye!")
	close(u.UserCh) // close the channel of the user
}

// user send a message to the server, broadcast the message to all online users of the server.
func (u *User) BroadcastMsg(msg string) {
	u.Server.BroadCast(u, ": " + msg)
}


func (u *User) HandleMsg(msg string) {
	if msg == "search" { // search online users of the server.
		u.Server.mapLock.Lock()
		for _, user := range u.Server.OnlineMap {
			onlineUser := "[" + user.Addr + "]" + user.Name + " is online"
			u.SendMessage(onlineUser) // send the online user to the client
		}
		u.Server.mapLock.Unlock()
	} else if strings.HasPrefix(msg, "rename|") {
		newName := strings.Split(msg, "|")[1]
		if newName == "" {
			u.SendMessage("rename failed, you are giving an empty name.")
		} else {
			u.Server.mapLock.Lock()
			_, ok := u.Server.OnlineMap[newName]
			if ok { // check if the new name is already taken by another user.
				u.SendMessage("rename failed, the name " + newName + " is already taken.")
				u.Server.mapLock.Unlock()
			} else {
				delete(u.Server.OnlineMap, u.Name)
				u.Server.OnlineMap[newName] = u
				u.Server.mapLock.Unlock()
				u.Name = newName
				u.SendMessage("rename success, the new name is " + newName)
			}
		}
	} else {
		u.BroadcastMsg(msg) // broadcast the message to all online users of the server.
	}
}
