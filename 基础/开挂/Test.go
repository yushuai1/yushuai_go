package main

import (
	"math/rand"
	"fmt"
)

/**
  模拟内存表
 */
func getArray(a,b int)([][]int)  {

	two :=make([][]int,512)

	for i:=0;i<a;a++ {
		intarray := make([]int,b)
		for j :=0;j<b;j++ {
			intarray[j] = rand.Intn(127)
		}
		two[i] = intarray
	}
	return two;
}


/**
  获取byte数组
 */
func getByte(lenth int) ([]byte) {

	b := make([]byte, lenth)
	for i := 0; i < lenth; i++ {
		b[i] = byte(rand.Intn(127))
	}
	return b;
}

func main() {

	b1 := getByte(512)
	b2 := getArray(12,12)


	fmt.Println(b2)
	fmt.Println(b1)
}
