package main

import (
	"sync"
	"fmt"
)

type singleton struct{
	name string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("只执行一次")
		instance = &singleton{}
	})
	return instance
}

func main() {

	my1 := GetInstance()
	my1.name = "uhyu"
	my2 :=GetInstance()
	fmt.Println(my2.name)
}
