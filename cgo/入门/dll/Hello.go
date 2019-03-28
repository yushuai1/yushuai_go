package main

import "C"
import (
"syscall"

	"flag"
	"fmt"
)

var(
	Path string
)
func callDll(){

	fmt.Println("--------------"+Path)
	//dll32 := syscall.NewLazyDLL("./TestDlls.dll")
	dll32 := syscall.NewLazyDLL(Path)
	println("call dll:",dll32.Name)
	g:=dll32.NewProc("fun")
	ret, _, _ :=g.Call(uintptr(4),uintptr(5))
	println()
	println("get the result:",ret)

}
func main() {
	flag.StringVar(&Path, "p", "p", "dll path")
	flag.Parse()
	callDll()
}