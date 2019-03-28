package main

import "fmt"

func Foo(param int)(n int, err error) {

	n = 1
	//m :=0
	//fmt.Println(n/m)
	return
}
func main() {

	s,e:=Foo(1)
	if e!=nil {
		fmt.Println(e)
	}else {
		fmt.Println(s)
	}
	fmt.Println("-------------")
}
