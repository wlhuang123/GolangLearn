package main

import (
	"bufio"
	"hwl/tool/logs"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://www.crazyflasher.com/skflashermsg/index.html")
	if err != nil {
		logs.Println(err)
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logs.Println("error http code:", resp.StatusCode)
	}

	r := bufio.NewReader(resp.Body)
	e := determainEcoding(r)

	utf8Reader := transform.NewReader(r, e.NewDecoder())

	content, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		logs.Println(err)
		return
	}
	logs.Println(string(content))
}

func determainEcoding(r *bufio.Reader) encoding.Encoding {
	content, err := r.Peek(1024)
	if err != nil {
		logs.Println(err)
		return unicode.UTF8
	}

	logs.Println(string(content))

	e, _, _ := charset.DetermineEncoding(content, "")
	return e

}
