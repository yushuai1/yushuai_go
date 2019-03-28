package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i  & 符号会生成一个指向其作用对象的指针
	fmt.Println(*p,p) // read i through the pointer
	*p = 21         // 类型 *p 是指向类型 p的值的指针。其零值是 nil 。
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j


	var s1 string

	s1 = "nihao"

	var pss *string

	pss = &s1

	fmt.Println(*pss)

}
