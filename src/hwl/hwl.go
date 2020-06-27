package main

import (
	"hwl/tool/logs"
	"io/ioutil"
	"net/http"
)

func main() {
	//learn.CopyTest()

	resp, err := http.Get("https://www.jianshu.com/p/ffc221e3e404")
	if err != nil {
		logs.Println(err)
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logs.Println("error http code:", resp.StatusCode)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Println(err)
		return
	}
	logs.Println(string(content))
}
