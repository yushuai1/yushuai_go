package main

import (
	"fmt"
	"runtime"
)

func xrange() chan int { // xrange用来生成自增的整数
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		for i := 0; ; i++ {
			ch <- i // 直到信道索要数据，才把i添加进信道
		}
	}()

	return ch
}

func cal() []int {
	generator := xrange()
	generator1 := xrange()

	yu := make([]int,10)
	for v:=0;v<10;v++ {
		yu[v] = <-generator+<-generator1
	}
	return yu
}

func main() {

	runtime.GOMAXPROCS(1)

	fmt.Println(cal())

}
