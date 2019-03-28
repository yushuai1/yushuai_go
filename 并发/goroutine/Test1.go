package main

import "fmt"


//信道用来goroutine之间通信  （int 类型就是chan int）
var complete chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	complete <- 0 // 执行完毕了，发个消息
}

/**
   加入go  相当于一个守护线程，主函数执行完成，就会推出并不会等go *
 */
func main() {

	go loop()
	//<- complete // 直到线程跑完, 取到消息. main在此阻塞住
	fmt.Printf("%d",<-complete)
}
