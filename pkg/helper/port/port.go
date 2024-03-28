package port

import "net"

//获取随机端口
func GenRandomPort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port
}
