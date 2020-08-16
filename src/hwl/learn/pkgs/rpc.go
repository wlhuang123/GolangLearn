package pkgs

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// RPCTest .
func RPCTest() {
	server()
	client()
}

// RPCExecFuncTest .
type RPCExecFuncTest struct{}

// Hello .
func (r *RPCExecFuncTest) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// 服务端代码
func server() {
	rpc.RegisterName("RPCExecFuncTest", new(RPCExecFuncTest))

	listener, err := net.Listen("tcp", ":9099")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

// 客户端代码
func client() {
	client, err := rpc.Dial("tcp", ":9099")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	var reply string
	if err := client.Call("RPCExecFuncTest.Hello", "req", &reply); err != nil {
		log.Fatal("Call error:", err)
	}

	fmt.Println("reply:", reply)
}
