package main

import (
	"sync"
	"fmt"
)
var MyLock  sync.Mutex


func LockTest(add *AddData)  bool{

	defer MyLock.Unlock()
	MyLock.Lock()
	for i:=0;i<10 ;i++  {
		fmt.Println("----",add.AppId)
	}

	return true
}



