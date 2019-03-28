package main

import (
	"fmt"
	"runtime"
)


func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}

//自定义错误类型
type ArithmeticErrorn struct {
	//cw string
}

//重写Error()方法
func (this *ArithmeticErrorn) Error() (s string) {
	s = "未知不确定异常，高危情况，需要立刻检查！！！"
	return
}

func ttt(num1, num2 int) (num int, er error) {

	defer func() {
		if r := recover(); r != nil {
			PrintStack()
			if se, ok := r.(error); ok {

				er = se
			} else {
				er = new(ArithmeticErrorn)
			}
		}
	}()

	if num2 == 0 {

		//panic(&ArithmeticErrorn{})
		num = num1 / num2
		//panic(&ArithmeticError1{}) //当然也可以使用ArithmeticError{}同时recover等到ArithmeticError类型
	} else {
		num = num1 / num2
	}

	return
}

func main() {

	s, f := ttt(1, 0)
	if f != nil {
		fmt.Println(f)
	} else {
		fmt.Println(s)
	}

	fmt.Println(1)

}
