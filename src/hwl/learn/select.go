package learn

import (
	"hwl/tool/logs"
	"time"
)

type selectTest struct{}

// SelectTest .
func SelectTest() {
	var s selectTest
	s.chanTimer()
	s.block()
	s.rand()
}

// rand 通过select随机选择生成数字
func (selectTest) rand() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 1:
			case ch <- 2:
			}
		}
	}()

	for v := range ch {
		logs.Println(v)
		time.Sleep(1 * time.Second)
	}
}

func (selectTest) block() {
	// select {} 如果还有goroutine在执行，该功能可以阻塞主函数不退出。但如果所有goroutine退出，程序会因为deaklock崩溃
}

// chanTimer 用select实现通道的超时判断
func (selectTest) chanTimer() {
	ch := make(chan int)

	go func() {
		for {
			select {
			case v := <-ch:
				logs.Println(v)
				return
			case <-time.After(2 * time.Second):
				logs.Println("time out")
				return
			}
		}
	}()

	go func() {
		time.Sleep(1 * time.Second) // 1s就上面的select能取出数据，3s就会超时
		ch <- 5
	}()

	time.Sleep(5 * time.Second) // 等待select里面取出chan数据并打印
}
