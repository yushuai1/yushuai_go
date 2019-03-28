package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	mytest:=make(map[interface{}]interface{})
	var lock sync.RWMutex

	var d interface{}
	t1 := time.Now().UnixNano()/1e6

	for i:=0;i<10000000;i++ {
		lock.Lock()
		mytest[i]=i
		d= mytest[i]
		lock.Unlock()
	}

	t2 := time.Now().UnixNano()/1e6

	fmt.Println(t2-t1,d)

}
