package learn

import (
	"fmt"
	"hwl/tool/logs"
)

/*
打印变量的方法
*/

type print struct {
}

// PrintTest .
func PrintTest() {
	var p print
	p.structValue().
		porinter().
		string().
		println().
		printCallString()
}

func (p *print) structValue() *print {
	type tmp struct {
		name string
	}

	var v tmp
	v.name = "abc"
	logs.Printf("%+v", v)
	return p
}

func (p *print) porinter() *print {
	var v *int
	logs.Printf("%p", v)
	return p
}

func (p *print) string() *print {
	s := "print learn"
	logs.Printf("%s", s)
	logs.Printf("%q", s) // %q打印字符串类型相比%s会增加双引号打印

	c := 'a'
	logs.Printf("%q", c) // %q还可以打印单引号的字符，%s打印会乱码

	return p
}

func (p *print) println() *print {
	s := "print learn"
	// %v是内置格式的任意值
	// Println就是默认用%v的方式来打印
	logs.Printf("%v", s)

	return p
}

type printCallT struct {
	x int
}

func (t printCallT) String() string {
	return "boo"
}

func (p *print) printCallString() {
	t := printCallT{123}
	fmt.Printf("%v\n", t)
}
