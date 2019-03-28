package main

/*
#include <stdio.h>
int hashcode_str(char* key) {
    unsigned int hashcode = 0;
    while (*key) {
        hashcode = (*key++) + (hashcode << 6) + (hashcode << 16) - hashcode;
    }
    hashcode &= 0x7FFFFFFF;
    printf("one is %d \n two is %d \n",hashcode,hashcode);
    return hashcode;
}
*/
import (
	"unsafe"
	"fmt"
)
import "C"


func main() {

	codeList:=make([]byte,3)
	codeList[1] = 1
	codeList[2] = 5
	codeList[0] = 11
	n:=C.hashcode_str((*C.uchar)(unsafe.Pointer(&codeList[0])))
	as:=n.(int32)
	fmt.Println(as)
}
