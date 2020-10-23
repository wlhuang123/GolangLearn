package learn

import (
	"fmt"
	"hwl/tool/logs"
	"sync"
	"time"
)

// ChannelTest .
func ChannelTest() {
	closeChanCanTrigerGetData()
	// readFromChannelWhenClose()
	// writeFromChannelWhenClose()
	// getCap()
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

// cap获取的是缓冲区的大小,和当前通道占用多少无关
// 获取当前占用使用len
func getCap() {
	c := make(chan int, 5)
	logs.Println("cat c:", cap(c), " len c:", len(c))

	c <- 1
	logs.Println("cat c:", cap(c), " len c:", len(c))
}

// closeChanCanTrigerGetData 关闭channel的时候，阻塞的chan可以读取出0值数据
func closeChanCanTrigerGetData() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ch:
				fmt.Println("ch get data")
				time.Sleep(1 * time.Second)
				fmt.Println("resource free")
				return
			default:
				fmt.Println("working")
			}
		}
	}()

	time.Sleep(4 * time.Second)
	close(ch)
	wg.Wait()
	fmt.Println("end")
}
