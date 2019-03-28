package main

import (
	"time"
	"fmt"
)

func main() {

	go spinner(1 * time.Microsecond)

	fmt.Println("***********************************")

	const n = 45

	fibN := fib(n)

	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {

	for i :=0;i<10000000;i++ {
		fmt.Printf("\r%s", "---------------")
		time.Sleep(delay)
	}

}

func fib(x int) int {

	fmt.Println("+++++++++++++++++++")
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
