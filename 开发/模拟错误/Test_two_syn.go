package main

import (
	"time"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(8)

	for k:=0;k<10;k++  {
		hj:=new(AddData)
		hj.AppId=k
		go LockTest(hj)
	}

	time.Sleep(time.Second*10)
}