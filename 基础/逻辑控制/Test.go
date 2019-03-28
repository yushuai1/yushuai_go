package main

import "fmt"

func main()  {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum :",sum)

	num :=0
	sum1 := 1
	for ; sum1 < 10;{
		num++
		fmt.Println("输出次数 :",num)
		sum1 += sum1
	}
	fmt.Println(sum1)

	sum2 := 1
	for sum2 < 1000 {
		sum2 ++
		//fmt.Println("输出次数: ",sum2)
	}
	fmt.Println(sum2)
}
