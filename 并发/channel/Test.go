package main

import "time"

func one(r chan <-int )  {


	r<-20
}

func two(v chan <-int )  {

	v<-2
}

func main() {

	var yu = make(chan int,2)

	defer func() {
		close(yu)
	}()
	yu<-1

	go one(yu)
	go two(yu)
	//two(yu)
	for tt := range yu{
		println(tt)
	}

	time.Sleep(time.Hour)

}
