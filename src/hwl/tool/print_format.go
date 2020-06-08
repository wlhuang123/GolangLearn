package tool

import "hwl/tool/logs"

// PrintFuncHeader .
func PrintFuncHeader(text string) {
	logs.Println("==", text)
}
