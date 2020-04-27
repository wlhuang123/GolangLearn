package pkgs

import (
	"hwl/tool/logs"
	"net/http"
	"os"
)

// NetHTTPTest .
func NetHTTPTest() {
	registServer()
	fileServer()
}

func fileServer() {
	http.Handle("/home/hwl", http.StripPrefix("/home/hwl", http.FileServer(http.Dir("/home/hwl"))))
}

func registServer() {
	http.HandleFunc("/helloworld", HelloWorld)
	http.HandleFunc("/index", Index)
	http.HandleFunc("/parseform", ParseForm)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

// ParseForm 解析参数
func ParseForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	logs.Println(name)

	name = r.FormValue("name")
	w.Write([]byte(name))
}

// Index 把文本的内容呈现返回，如果是html就能呈现html的界面格式
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")

	f, _ := os.Open("/home/hwl/Code/GolangLearn/src/learn/pkg/index.js")
	defer f.Close()

	// 1.直接拷贝方式
	// io.Copy(w, f)
	// 2.读取文件byte内容方式
	// buf, _ := ioutil.ReadAll(f)
	// w.Write(buf)
	// 3.f.Read 读取到容器
	buf := make([]byte, 1024*10)
	ln, _ := f.Read(buf)
	w.Write(buf[:ln])
}

// HelloWorld ,
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
