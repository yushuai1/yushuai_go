package main

import (
	"fmt"
	"sync"
	"time"
)
type structR6 struct {
	B1 [10]int
	Name string
}

var r6Pool = sync.Pool{
	New: func() interface{} {
		return new(structR6)
	},
}
func usePool() {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		sr6 := r6Pool.Get().(*structR6)
		sr6.B1[0] = 0
		r6Pool.Put(sr6)
	}
	fmt.Println("pool Used:", time.Since(startTime))
}
func standard() {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		var sr6 structR6
		sr6.B1[0] = 0
	}
	fmt.Println("standard Used:", time.Since(startTime))
}
func main() {
	//standard()
	//usePool()

	sr6 := r6Pool.Get().(*structR6)
	sr6.Name = "nihao"
	r6Pool.Put(sr6)
	fmt.Println(sr6)
	sr7 := r6Pool.Get().(*structR6)
	fmt.Println(sr7)
}
