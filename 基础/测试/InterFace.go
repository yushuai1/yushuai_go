package main


/**
  接口的特性好处
  1：能够内部自动传递结构体的值
  2：完整性约束
 */

import (
	"fmt"
	"yushuai/基础/函数"
	"yushuai/基础/接口"
	"yushuai/基础"
)

func main() {


	fmt.Println(iui.Name)

	var ch 函数.People

	chs:=&接口.Chinese{Name:"asdasd",Energy:655}
	ch = chs
	ch.Work(10)
	ch.Sleep(10)
	fmt.Println("chinese",*chs)

	var eg 函数.People
	eg = &接口.English{Name:"asdasd",Age:655}
	eg.Work(10)
	eg.Sleep(10)
	fmt.Println("English")
}
