package main



/*
#include<stdlib.h>
#include <stdio.h>
int cal(__int8 *a,__int8 *b){
    int len1=*a;
    int len2=*b;
    printf("one is %d \n two is %d \n",len1,len2);
    return len1+len2;

}
*/
import "C"
import "fmt"
import (
	"unsafe"
	"reflect"
)

func main() {
	a:=make([]int,2)
	b:=make([]int,2)

	a[0] = 100
	b[0] = 101

	//aa:= &a
	//bb :=&b

	nk:=C.cal((*C.cahr)(unsafe.Pointer(&a[0])),(*C.int)(unsafe.Pointer(&b[0])))

	fmt.Println(reflect.TypeOf(nk),"-----------")
}