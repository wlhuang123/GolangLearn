package pkgs

import (
	"hwl/tool/logs"
	"time"
)

/*
  时间使用相关的
*/

// TimeTest .
func TimeTest() {
	calculateUseTime()
}

// calculateUseTime 统计耗时
func calculateUseTime() {
	begin := time.Now()
	time.Sleep(1 * time.Second)
	end := time.Now()

	logs.Println(end.Sub(begin))
}
