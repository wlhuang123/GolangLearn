package pkgs

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

var count int

// NetTCPClientTest .
func NetTCPClientTest() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9998")

	for {
		count++
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Println("dia error:", err)
			time.Sleep(1 * time.Second)
		}
		defer conn.Close()

		fmt.Println("connected! count:", count)
		onMessageRecived(conn, count)
	}
}

// NetTCPServerTest .
func NetTCPServerTest() {
	var count int

	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9998")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("accetp error:", err)
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		count++
		go tcpPipe(tcpConn)
	}
}

func onMessageRecived(conn *net.TCPConn, count int) {
	reader := bufio.NewReader(conn)

	fmt.Println("new reader sucess")

	b := strconv.Itoa(count)
	//fmt.Scanln(&b)
	_, err := conn.Write([]byte("client count: " + b + "\n"))
	if err != nil {
		fmt.Println("write error:", err)
	} else {
		fmt.Println("write sucess")
	}

	msg, err := reader.ReadString('\n')
	fmt.Println(msg)
	if err != nil {
		fmt.Println("read error:", err)
	} else {
		fmt.Println("read sucess msg:", msg)
	}
	time.Sleep(time.Microsecond)

}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	fmt.Println("new reader sucess")

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read error", err)
			return
		}

		fmt.Println("read sucess:", string(message))
		countb, err := conn.Write([]byte("client source return:" + string(message) + "\n"))
		fmt.Println("is write sucess err:", err, " countb:", countb, " count:", count)
	}
}
