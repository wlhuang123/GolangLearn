package learn

import (
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
		porinter()
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
