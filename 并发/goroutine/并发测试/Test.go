package main

import (
	"fmt"
	"sync"

)

func xrange() <-chan int { // xrange用来生成自增的整数
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		defer func() {
			close(ch)
			fmt.Println("处理程序1也要退出！")
		}()
		for i := 0; i<1000; i++ {
			ch <- i // 直到信道索要数据，才把i添加进信道
		}
	}()

	return ch
}

/**
   chan 用完就要关闭，为了防止出现异常，最好在生产数据端关闭
 */
func 产生数据(done <-chan struct{},num <-chan int) <-chan int {

	数据产生结果chan := make(chan int, 100)
	go func() {
		defer func() {
			close(数据产生结果chan)
			fmt.Println("处理程序1也要退出！")
		}()
		for values := range num {

			select {
			case 数据产生结果chan <- (values * 10):
			case <- done:
				return
			}
		}

	}()

	return 数据产生结果chan;
}

func 处理数据_1(done <-chan struct{}, ch <-chan int) <-chan int {

	数据处理结果chan := make(chan int, 100)
	go func() {
		defer func() {
			close(数据处理结果chan)
			fmt.Println("处理程序也要退出！")
		}()
		for n := range ch {
			select {
			case 数据处理结果chan <- (n + 555):
			case <-done:
				fmt.Println("主线程退出了 处理程序也要退出！")
				return
			}
		}
	}()

	return 数据处理结果chan
}

/**
  为了提高cpu和io的利用率，我们需要让对个程序对应一个chan，也就是增大扇出
 */
func 合并chan(done <-chan struct{}, cs ... <-chan int) <-chan int {

	var wg sync.WaitGroup

	合并结果chan := make(chan int, 100)

	合并结果method := func(c <-chan int) {

		for nk := range c {
			select {
			case 合并结果chan <- nk:
			case <-done:
				fmt.Println("主线程退出了 合并程序也要退出！")
				return
			}

		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go 合并结果method(c)
	}

	go func() {
		wg.Wait()
		close(合并结果chan)
	}()
	return 合并结果chan

}

func test() int{

	数据 := xrange()
	done := make(chan struct{})

	defer func() {
		close(done)
		//time.Sleep(time.Hour)
		fmt.Println("主线程退出")
	}()

	第一步处理_1 := 产生数据(done,数据)
	第一步处理_2 := 产生数据(done,数据)

	第二步处理_1 := 处理数据_1(done, 第一步处理_1)
	第二步处理_2 := 处理数据_1(done, 第一步处理_2)
	合并 := 合并chan(done, 第二步处理_1, 第二步处理_2)


	sum:=0

	//for v:=0;v<10;v++ {
	//	sum = sum+<-合并
	//}
	for k := range 合并 {
	    sum = sum+k
	}
	fmt.Println(sum)

	return sum

}

func main()  {

   fmt.Println(test(),"------------")

}
