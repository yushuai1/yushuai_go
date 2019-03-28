package main

import (
	"time"
	"sync"
	"fmt"
)

func main() {
	var mapTest sync.Map

	var d interface{}
	t1 := time.Now().UnixNano()/1e6

	for i:=0;i<10000000;i++ {
		mapTest.Store(i,i)
		d,_ = mapTest.Load(i)
	}

	t2 := time.Now().UnixNano()/1e6

	fmt.Println(t2-t1,d)
}
