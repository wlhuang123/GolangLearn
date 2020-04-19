package learn

import "hwl/tool/logs"

/*变量声明*/

// VarDeclareTest .
func VarDeclareTest() {
	// :=用于短变量的声明方式 简洁方便，但是全局变量不能这样声明
	s := "var declare"
	logs.Println(s)

	// 在局部和全局都可以使用这种声明方式，并且会初始化为内置类型的默认值
	var c string
	logs.Println(c)

	// 很少使用这种方式，除非声明多个变量
	var d = ""
	logs.Println(d)

	// 很少使用。在赋值类型一致的情况下属于冗余的声明方式。 在类型不一致的时候声明是必需的。
	// 也为了说明初始化声明的变量类型的重要性
	var e uint16 = 1
	logs.Println(e)

}
