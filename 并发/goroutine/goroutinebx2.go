package main

import (
	"fmt"
	"runtime"
)

var quits chan int = make(chan int)

func loops() {
	for i := 0; i < 10; i++ {
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine,人眼看上去并行了
		fmt.Printf("%d ", i)
	}
	quits <- 0
}


func main() {

	go loops()
	go loops()

	for i := 0; i < 2; i++ {
		<- quits
	}
}