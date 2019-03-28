package main

import (
	"syscall"
	"unsafe"
	"fmt"
)

func main() {

	byt := make([]byte,2)
	byt[0] = 12


	a:=make([]int,2)
	b:=make([]int,2)

	a[0] = -100
	b[0] = -101


	fmt.Println(&a[0],&b[0])
	fmt.Println((uintptr)(unsafe.Pointer(&a[0])))
	dll := syscall.NewLazyDLL("./AK.dll")

	cal:=dll.NewProc("cal")

	sy:=dll.NewProc("sy")

	ret, _, _ :=cal.Call((uintptr)(unsafe.Pointer(&a[0])),(uintptr)(unsafe.Pointer(&b[0])))
	retsy, _, _ :=sy.Call((uintptr)(unsafe.Pointer(&byt[0])),uintptr(5))

	println("get the result:",ret)
	println("get the result:",retsy)
}
