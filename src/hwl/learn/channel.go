package learn

import "hwl/tool/logs"

// ChannelTest .
func ChannelTest() {
	readFromChannelWhenClose()
	writeFromChannelWhenClose()
}

func readFromChannelWhenClose() {
	c := make(chan int, 10)
	c <- 1
	close(c)

	a, ok := <-c
	logs.Println("read from channel when channle is closed, still get before send data:", a, ok)

	a, ok = <-c
	logs.Println("read from channel when channle is closed, and channel no data , will get zero data:", a, ok)
}

// 给已经关闭的channel发送数据会panic
func writeFromChannelWhenClose() {
	defer func() {
		err := recover()
		logs.Println(err)
	}()
	c := make(chan int, 10)
	close(c)
	c <- 1
}
