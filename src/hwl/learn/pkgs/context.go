package pkgs

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ContextTest .
func ContextTest() {
	timeoutContext()
}

// timeoutContext 带超时的context
func timeoutContext() {
	work := func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done(): // 如果执行cancel()，这里就会接收到chan返回数据
				fmt.Println(ctx.Err()) // 打印出取消原因：context canceled
				return
			default:
				fmt.Println("working")
			}
		}
	}

	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(ctx, wg)
	}

	time.Sleep(6 * time.Second)
	cancel()
	wg.Wait()
}
