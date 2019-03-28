package main

/*
#include <stdio.h>
#include <TestDll.h>
*/
import "C"
import (
	"fmt"
)

func main() {
	cstr := C.times(2,5)
	fmt.Println(cstr)
}