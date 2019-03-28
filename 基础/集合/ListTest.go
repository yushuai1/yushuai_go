package main

import ("fmt"
	"container/list"
	"reflect"
)

func main()  {

	listOne:=list.New()

	fmt.Println(reflect.TypeOf(listOne).String(),"--------------------------------")

	listOne.PushBack("nihao")
	listOne.PushBack(545)
	listOne.PushBack(true)

	fmt.Println(listOne.Len())

	for v:=listOne.Front();v!=nil;v = v.Next() {
		fmt.Println(v.Value)
	}

	r:= listOne.PushFront(2)
	fmt.Println(r,"--------------------------")

	for v:=listOne.Front();v!=nil;v = v.Next() {
		fmt.Println(v.Value)
	}
}