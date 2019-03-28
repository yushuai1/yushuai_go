package main

import (

	"github.com/goinggo/mapstructure"
	"fmt"
)

//这里对应的 N 和 A 不能为小写，首字母必须为大写，这样才可对外提供访问，具体 json 匹配是通过后面的 tag 标签进行匹配的，
// 与 N 和 A 没有关系
//tag 标签中 json 后面跟着的是字段名称，都是字符串类型，要求必须加上双引号，否则 golang 是无法识别它的类型
type Person struct {
	Name string     `mapstructure:"name"`
	AgeMe int        `mapstructure:"age_me"`
}

func TestMap2Struct() {
	mapInstance := make(map[string]interface{})
	mapInstance["name"] = "liang637210"
	mapInstance["age_me"] = 28

	var person Person
	//将 map 转换为指定的结构体
	if err := mapstructure.Decode(mapInstance, &person); err != nil {
		fmt.Println(err)
	}
	fmt.Println( person)
}

func main() {
	TestMap2Struct()
}
