package leetcode

import (
	"fmt"
	"io/ioutil"
)

// GenBasicCode 生成一个leecode题目的基本代码
// num是题目序号 eg:023
// name是题目名字 eg:three_num_sum
func GenBasicCode(num, name string) {
	workdir := "/home/hwl/Code/GolangLearn/src/hwl/leetcode/"
	file := workdir + num + "_" + name + ".go"

	var code string

	code += "package leetcode\n\n"
	code += GenFunc("Test"+num) + "\n\n" + GenFunc("Check"+num)

	ioutil.WriteFile(file, []byte(code), 0777)
}

// GenFunc .
func GenFunc(funcName string) string {
	return fmt.Sprintf("// %s .\nfunc %s() {\n\n}", funcName, funcName)
}
