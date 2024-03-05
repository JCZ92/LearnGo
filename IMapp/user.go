package main
import (
	"net"
)

type User struct {
	Name string
	Addr string
	UserCh chan string
	conn net.Conn
} 

// factory for a new User instance
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		UserCh: make(chan string), // unbuffered channel
		conn: conn,
	}
	go user.PushMessageToClient() // start a goroutine to push message to the client
	
	return user
}

// get the message from the channel and send it to the client
func (u *User) PushMessageToClient() {
	for msg := range u.UserCh {
		u.conn.Write([]byte(msg + "\n"))
	}
}