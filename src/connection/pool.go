package connection

import (
	"net"
)

// Config the connect config
type Config struct {
	// the heartbeat interval
	interval int32
	// transcat timeout
	timeout int32
}

// Connection : a valid connection
type Connection struct {
	// record the connection status
	// include the packet success rate
	status map[int32]int32
	// the connection config
	config Config
	// ip
	ipaddr string
	//
	// conn net.Conn
	//
}

func (con *Connection) entry() {
	conn, error := net.Dial("tcp", con.ipaddr)
	if error != nil {
		return
	}
	conn.Close()

}
