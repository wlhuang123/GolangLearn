package logs

import (
	"log"
	"os"
)

var f *log.Logger

func init() {
	f = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	if f == nil {
		log.Println("f==nil")
		return
	}
}

// Println .
func Println(v ...interface{}) {
	f.Println(v)
}

// Printf .
func Printf(format string, v ...interface{}) {
	f.Printf(format, v)
}
