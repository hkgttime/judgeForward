package channel

import (
	"net"
	"onlinejudgeForward/slflog"
)

var addr string = "127.0.0.1:8000"

func sender(addr, body string)  {
	conn, err := net.Dial("tcp", addr)
	slflog.FatalErr(err, "conn err")
	defer conn.Close()
	_, err = conn.Write([]byte(body))
	slflog.FatalErr(err, "write err")
}

func reader(conn net.Conn)  {
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	slflog.FatalErr(err, "read err")
	slflog.Debug(string(buf[:len]))

}


