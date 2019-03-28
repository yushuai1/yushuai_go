package main

import (
	"fmt"
)


func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main()  {

	var slice1 = make([]int ,5,10)
	fmt.Println(slice1)

	s := []int {1,2,3}
	b :=append(s, 4)
	s[0] = 100
	fmt.Println(s,b)

	n :=slice1[:]
	fmt.Println(n)




	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)


	fmt.Println("-----------------------------------------------appand-----------------------------------------------")



	var k,j []int
	printSlice(k)
	k = make([]int,2)
	printSlice(k)
	fmt.Println(cap(k),len(k),k)

   fmt.Println("-------------------")
	j=append(k,1)
	j[0] = 100
	fmt.Println(cap(j),len(j),j)
	fmt.Println(cap(k),len(k),k)

	fmt.Println("**********************")

	vc :=make([]int,2,10)
	fmt.Println(cap(vc),len(vc),vc)
	vb:=append(vc,2)
	fmt.Println(cap(vb),len(vb),vb)
	vc[0]=101
	fmt.Println(cap(vb),len(vb),vb)
	fmt.Println(cap(vc),len(vc),vc)


	fmt.Println("-----------------*************************------------------------")

	slice11 := []byte{1, 2, 3, 4, 5}
	slice22 := []byte{5, 4, 3}

	copy(slice11, slice11) // 只会复制slice1的前3个元素到slice2中
	fmt.Printf("%x\n",slice11)
	fmt.Printf("%x\n",slice22)

	for k:=0;k<3;k++ {
		slice11[k] =slice22[k]
	}

	fmt.Println(slice11)

	fmt.Println("11111111111111111111111111111111111111111111111111111111111111")
	var kas = make([]int,0,3)
	kaas :=append(kas, 4)
	kas = kaas

	fmt.Println(kas)





}