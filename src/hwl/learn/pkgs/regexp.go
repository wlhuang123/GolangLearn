package pkgs

import (
	"hwl/tool/logs"
	"regexp"
)

// RegexpTest .
func RegexpTest() {
	SimpleChar()
}

// SimpleChar 简单基本的匹配
// 下面的例子都是match的
func SimpleChar() {
	// 匹配到0-6以内才认为通过，同理（a-z,A-Z，但是不能写成a-B不知道会识别成什么）
	// []匹配单个字符
	// ^和$用来指明开头和结尾
	ShowMatchStatus("^[0-6]$", "6")

	// ^反向字符族，就是非的意思。匹配除了9，d以外的任意单个字符
	ShowMatchStatus("^[^9d]$", "o")

	// 匹配数字或小写字母。 +是重复一次或更多次
	ShowMatchStatus("^[a-z0-9]+$", "ahjk8")

	// 匹配数字或小写字母。 ?是重复0次或1次
	ShowMatchStatus("^[a-z0-9]?$", "9")

	// 匹配数字或小写字母。 *是重复任意次
	ShowMatchStatus("^[a-z0-9]*$", "9sfa")

	// 只要包含a78的字串即匹配
	ShowMatchStatus("a78", "0a78k@")

	// *匹配0个或多个字符
	ShowMatchStatus("a7*8", "0a78k@")

	// 对于 [a-z] 这样的正则表达式，如果要在 [] 中匹配 - ，可以将 - 放在 [] 的开头或结尾，例如 [-a-z] 或 [a-z-]
	ShowMatchStatus("[-a-z]", "-")

	// 长度为3-5的数字或小写字母,{}用来指明重复次数，如果里面只有一个数字，则是固定次数。{n,}重复n或更多次
	ShowMatchStatus("^[a-z0-9]{3,5}$", "449p")

	// 长度为3-5的数字或小写字母,还可以带特殊字符@
	ShowMatchStatus("^[a-z0-9@]{3,5}$", "44p@")

	// 长度11位的，1开头的手机号码
	ShowMatchStatus("^1[0-9]{10}$", "17322343615")

	// 长度11位的，1开头的手机号码。 \d匹配数字
	ShowMatchStatus(`^1[\d]{10}$`, "17322343615")

	// 小数 匹配"." "("需要加转义  如果用"",需要\\转义 "\\d+\\.[0-9]+$"
	ShowMatchStatus(`\d+\.[0-9]+$`, "0.9")

	// 邮箱
	ShowMatchStatus(`^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`, "80@qq.com")
}

// ShowMatchStatus .
func ShowMatchStatus(pattern string, s string) {
	ok, _ := regexp.MatchString(pattern, s)
	if ok {
		logs.Println("Match: ", pattern, " ", s)
	} else {
		logs.Println("not Match: ", pattern, " ", s)
	}
}
