package main

import (
	"crawl/mycase"
	"hwl/tool/logs"
)

func main() {
	url := "http://www.crazyflasher.com/skflashermsg/index.html"
	content := mycase.GetURLContent(url)
	logs.Println(content)
}
