package main

import "fmt"

func isChanClose(ch chan int) bool {
	select {
	case _, ok := <-ch:
		return !ok
	default:
	}
	return false
}

func main() {
	ch := make(chan int, 4)

	close(ch)

	fmt.Println(isChanClose(ch))
}
