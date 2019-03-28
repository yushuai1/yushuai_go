package main

import "fmt"

// 表示声明的变量的初始类型，当没有int时候自动根据赋值变成相应的类型
var i,j int = 1,3

const PI  = 425

func main()  {
	var c,python,java = true,false,"no!"

	//在函数中， := 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
	//函数外的每个语句都必须以关键字开始( var 、 func 、等等)， :=结构不能使用在函数外。
	k := 3
	fmt.Println(i,j,c,python,java,k)


	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s,PI)
}
