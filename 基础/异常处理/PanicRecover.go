package main

import (
	"fmt"
)

//自定义错误类型
type ArithmeticError1 struct {
	error
}

//重写Error()方法
func (this *ArithmeticError1) Error() string {
	return "自定义的error,error名称为算数不合法"
}

//定义除法运算函数***这里与本文中第一幕①error接口的例子不同
func Devide1(num1, num2 int) int {
	if num2 == 0 {
		panic(&ArithmeticError1{}) //当然也可以使用ArithmeticError{}同时recover等到ArithmeticError类型
	} else {
		return num1 / num2
	}
}
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic的内容%v\n", r)
			if _, ok := r.(error); ok {
				fmt.Println("panic--recover()得到的是error类型")
			}
			if _, ok := r.(*ArithmeticError1); ok {
				fmt.Println("panic--recover()得到的是ArithmeticError类型")
			}
			if _, ok := r.(string); ok {
				fmt.Println("panic--recover()得到的是string类型")
			}
		}
	}()

	rs := Devide1(a, b)
	fmt.Println("结果是：", rs)
}