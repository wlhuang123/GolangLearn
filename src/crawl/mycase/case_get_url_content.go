package mycase

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

/*
   读取一个网页的内容
*/

// GetURLContent .
func GetURLContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		logs.Println(err)
		return ""
	} else if resp.StatusCode != http.StatusOK {
		logs.Println("error http code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()

	r := bufio.NewReader(resp.Body)
	e := DetermainEcoding(r)

	readerEncoding := transform.NewReader(r, e.NewDecoder())

	// 这里读取内容不能直接读取resp.Body，网页编码utf-8是没问题的，但是有的中文网是GBK的编码，直接读取就会乱码
	// 所以需要先获取网页的编码，再用该编码来读取内容
	content, err := ioutil.ReadAll(readerEncoding)
	if err != nil {
		logs.Println(err)
		return ""
	}
	return string(content)
}

// DetermainEcoding .
func DetermainEcoding(r *bufio.Reader) encoding.Encoding {
	content, err := r.Peek(1024)
	if err != nil {
		logs.Println(err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(content, "")
	return e
}
