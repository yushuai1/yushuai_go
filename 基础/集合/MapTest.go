package main

import ("fmt")

func main()  {

	//创建
	var mapOne  map[string]string
	fmt.Println(mapOne)
	mapOne =make(map[string]string)

	//mapOne:=make(map[string]string)

	//添加
	mapOne["ni"] = "hao"
	mapOne["wo"] = "buhao"
	mapOne["what"] = "nimei"
	fmt.Println(len(mapOne))
	fmt.Println(mapOne)

	//遍历
	for key := range mapOne{
		fmt.Println("key is = ",key,"   values is = ",mapOne[key])
	}

	//删除
	delete(mapOne, "ni")

	for key := range mapOne{
		fmt.Println("key is = ",key,"   values is = ",mapOne[key])
	}
}