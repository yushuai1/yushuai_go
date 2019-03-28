package main


/*
#include <stdio.h>

int hashcode_str(unsigned char* key) {
    unsigned int hashcode = 0;
    while (*key) {
        hashcode = (*key++) + (hashcode << 6) + (hashcode << 16) - hashcode;
    }
    hashcode &= 0x7FFFFFFF;
    //printf("one is %d \n two is %d \n",hashcode,hashcode);
    return hashcode;
}
*/
import "C"
import (
	"fmt"
	"time"
)

func main() {
	var n int
	t1 := time.Now().UnixNano()/1e6
	for mn:=100;mn<100000000 ;mn++  {
		n = mn%63
		//b:=[]byte("yushuai_"+strconv.Itoa(mn))
		//n=int(C.hashcode_str((*C.uchar)(unsafe.Pointer(&b[0]))))
	}
	t2 := time.Now().UnixNano()/1e6
	fmt.Println(n,"------------",(t2-t1))

}