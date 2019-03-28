package main

import (
	"fmt"
	"runtime"
)

var quitb chan int = make(chan int)

func loopb() {
	for i := 0; i < 100; i++ { //为了观察，跑多些
		fmt.Printf("%d ", i)
	}
	quitb <- 0
}


//这才是真正的并行，因为默认情况下
//runtime.GOMAXPROCS（1）也就是一个核心，一旦程序跑起来只要自己不睡眠根本不会让出cpu
func main() {
	runtime.GOMAXPROCS(2) // 最多使用2个核

	go loopb()
	go loopb()

	for i := 0; i < 2; i++ {
		<- quitb
	}
}