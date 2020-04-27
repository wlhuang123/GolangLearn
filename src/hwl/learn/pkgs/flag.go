package pkgs

import (
	"flag"
	"fmt"
	"hwl/tool/logs"
	"time"
)

// FlagTest .
func FlagTest() {
	durationTest()
	implementMyValue()
}

type myValue struct {
	name string
}

func (m *myValue) String() string {
	return "i don't know usage"
}

// Set 在flag.Parse()里面会自动调用这个方法，然后就能把输入的s赋值给m.name
func (m *myValue) Set(s string) error {
	fmt.Sscanf(s, "%s", &m.name)
	return nil
}

// implementMyValue 自定义类型使用flag
func implementMyValue() {
	/* 自定义的类型需要实现下面的接口
		type Value interface {
				String() string
				Set(string) error
	    }
	*/
	f := myValue{name: "hwl"}
	flag.CommandLine.Var(&f, "myValue", "test")

	flag.Parse()
	logs.Println(f.name)
}

func durationTest() {
	// args1 命令行使用的参数名,  args2 默认值, args3 描述
	// 返回输入的值
	// 使用 ./proc -period 1h5m30s
	// ./proc -h就会默认打印使用说明
	period := flag.Duration("period", 1*time.Second, "sleep period")
	time.Sleep(*period)
}
