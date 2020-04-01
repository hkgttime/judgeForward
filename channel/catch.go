package channel

import (
	"net"
	"onlinejudgeForward/slflog"
)

type status uint8

var judgeList []JudgeInfo

func init() {
	judgeList = make([]JudgeInfo, 5)
}

func GetJudgeList() []JudgeInfo {
	return judgeList
}

func StartCatch()  {
	catch("127.0.0.1:5588")
}

type JudgeInfo struct {
	Host string
	Port int
	Status status
	ProcessNum int
	OperatingLoad float32
}

func catch(addr string)  {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	slflog.FatalErr(err, "ResolveUDPAddr err")
	udp, err := net.ListenUDP("udp", udpAddr)
	slflog.FatalErr(err, "ListenUDP err")
	defer udp.Close()
	for {
		buf := make([]byte, 1024)
		len, raddr, err := udp.ReadFromUDP(buf)
		slflog.FatalErr(err, "ReadFromUDP err")
		slflog.Debug(string(buf[:len]))
		_, err = udp.WriteTo([]byte("ack"), raddr)
		judge, errs := parseUdp(string(buf[:len]))
		slflog.FatalErr(errs, "")
		judgeList = append(judgeList, *judge)
		slflog.FatalErr(err, "WriteTo err")
	}

}

func parseUdp(data string) (*JudgeInfo, error) {
	judge := new(JudgeInfo)
	return judge, nil
}
