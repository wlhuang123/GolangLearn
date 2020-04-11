package learn

import (
	"hwl/tool/logs"
	"sync"
	"sync/atomic"
)

/*
   golang封装的原子操作函数的使用
*/

// ActomicTest .
func ActomicTest() {
	noAtomicAdd()
	atomicAdd()
	atomicSetAndLoad()
}

// noAtomicAdd 没有原子操作函数并发对一个变量自增会产生意外结果
func noAtomicAdd() {
	var wg sync.WaitGroup
	wg.Add(100000)

	var result int

	for i := 0; i < 100000; i++ {
		go func() {
			result++
			wg.Done()
		}()
	}

	wg.Wait()
	// result自增10万次,但是结果大概率不是10万
	// result++不是原子操作,有可能前面加1了,后面的goroutine在加1前的结果加1再保存回来,就少加了一次
	logs.Println("noActomicAdd, expect 100000, but result:", result)
}

func atomicAdd() {
	var wg sync.WaitGroup
	wg.Add(100000)

	var result int64

	for i := 0; i < 100000; i++ {
		go func() {
			atomic.AddInt64(&result, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	// 前面用了atomic.AddInt64(&result, 1) 自增,所以结果是10万
	logs.Println("atomicAdd, expect 100000,  result:", result)
}

func atomicSetAndLoad() {
	var shutdown int64
	atomic.StoreInt64(&shutdown, 1)       // 安全的设置值
	if atomic.LoadInt64(&shutdown) == 1 { // 安全的取出值判断  可用于在另一个goroutine里面判断是否流程结束关闭

	}
}
