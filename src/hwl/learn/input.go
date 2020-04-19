package learn

import (
	"bufio"
	"hwl/tool/logs"
	"os"
)

/*
   对输入进行处理的方法，输入包含：文件输入，Stdin输入
*/

// InputLearnTest .
func InputLearnTest() {
	bufioTest()
}

// bufioTest 逐行读的方法
func bufioTest() {
	// 1.文件的逐行读取
	f, err := os.Open("/home/hwl/Code/GolangLearn/hwl.sh")
	if err != nil {
		logs.Println(err)
		return
	}
	defer f.Close()
	inputFile := bufio.NewScanner(f)
	// 可以用for循环对input逐行获取内容
	for inputFile.Scan() {
		logs.Println(inputFile.Text())
	}

	// 2.Stdin的逐行读取
	inputStdin := bufio.NewScanner(os.Stdin)
	// 可以用for循环对input逐行获取内容
	// Scan在读取到新行的时候返回true，Text获取读到的内容
	for inputStdin.Scan() {
		// 这里可以获取到一行的内容，可以利用map，把行内容当成key，达到去除重复行的算法
		logs.Println(inputStdin.Text())
	}

}
