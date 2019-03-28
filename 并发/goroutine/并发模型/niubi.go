package main

import (
	"fmt"
	"sync"
)


func merge(done <-chan struct{},cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int,1)
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

//func gen(nums ...int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for _, n := range nums {
//			out <- n
//		}
//		close(out)
//	}()
//	return out
//}


func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}


func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

//func main() {
//	// Set up the pipeline.
//	c := gen(2, 3)
//	out := sq(c)
//	// Consume the output.
//	fmt.Println(<-out) // 4
//	fmt.Println(<-out) // 9
//}

//func main() {
//	for n := range sq(sq(gen(2, 3))) {
//		fmt.Println(n) // 16 then 81
//	}
//}

//func main() {
//	in := gen(2, 3)
//
//	// Distribute the sq work across two goroutines that both read from in.
//	c1 := sq(in)
//	c2 := sq(in)
//
//	// Consume the merged output from c1 and c2.
//	for n := range merge(c1, c2) {
//		fmt.Println(n) // 4 then 9, or 9 then 4
//	}
//}


func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.


	// Consume the first value from output.
	done := make(chan struct{}, 2)
	c1 := sq(done,in)
	c2 := sq(done,in)
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Tell the remaining senders we're leaving.
	done <- struct{}{}
	done <- struct{}{}
}