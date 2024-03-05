package main
import (
	"net"
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
// user come online, add the user to the online user map of the server,
// and broadcast the user online status to all online users of the server.
func (u *User) Online() {
		// add user to the online user map of the server
		u.Server.mapLock.Lock()
		u.Server.OnlineMap[u.Name] = u
		u.Server.mapLock.Unlock()
		// broadcast the user online status to all online users of the server
		u.Server.BroadCast(u, "is online now. Welcome!")
}

// user go offline, remove the user from the online user map of the server,
// and broadcast the user offline status to all online users of the server.
func (u *User) Offline() {
	// remove user from the online user map of the server
	u.Server.mapLock.Lock()
	delete(u.Server.OnlineMap,u.Name)
	u.Server.mapLock.Unlock()
	// broadcast the user online status to all online users of the server
	u.Server.BroadCast(u, "is offline. Bye!")
}

// user send a message to the server, broadcast the message to all online users of the server.
func (u *User) BroadcastMsg(msg string) {
	u.Server.BroadCast(u, msg)
}


func (u *User) HandleMsg(msg string) {
	if msg == "search" { // search online users of the server.
		u.Server.mapLock.Lock()
		for _, user := range u.Server.OnlineMap {
			onlineUser := "User [" + user.Addr + "]" + user.Name + ": is online"
			u.UserCh <- onlineUser // send the online user to the client
		}
		u.Server.mapLock.Unlock()
	} else {
		u.BroadcastMsg(msg) // broadcast the message to all online users of the server.
	}

}