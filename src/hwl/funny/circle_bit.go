package funny

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 蒙纳卡洛 计算圆周率

// CircleBitTest .
func CircleBitTest() {
	var PI float64
	var testNum uint64 = 100000000

	var circleNum uint64

	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)

	for i := 0; i < int(testNum); i++ {
		randomx := r1.Float64() * 2
		randomy := r1.Float64() * 2
		if isInCircle(randomx, randomy) {
			circleNum++
		}
	}

	PI = 4 * float64(circleNum) / float64(testNum)
	fmt.Println(circleNum, testNum)
	fmt.Println("PI:", PI)
}

func isInCircle(x, y float64) bool {
	lenx := math.Abs(x - 1)
	leny := math.Abs(y - 1)
	result := math.Sqrt(lenx*lenx + leny*leny)
	if result <= 1 {
		return true
	}
	return false
}
