package main

import (
	"hwl/tool"
	"hwl/tool/logs"
)

func main() {
	bit := tool.GetHostBits()
	logs.Println(bit)
}
