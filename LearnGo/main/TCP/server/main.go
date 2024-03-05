package main
import (
	"fmt"
	"net"
)
func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("read success:", string(buf[:n]))
	}	
}

func main() {
	ls, err := net.Listen("tcp", "localhost:8999")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	for {
		conn, err2 := ls.Accept()
		if err2 != nil {
			fmt.Println("accept err:", err2)
			return
		} else {
			fmt.Println("accept success: client =", conn.RemoteAddr().String())
		}
		go process(conn)
	}
	
}