package main

import (
	"fmt"
	//"yushuai/基础/函数"
)

type Vertex struct {
	X int
	Y int
	Z string
}

/**
   相当于new一个新的对象
 */
func (v *Vertex) getx(x int) (Vertex) {
	vb := Vertex{v.X, v.Y + x, v.Z}
	return vb
}

func (v *Vertex) getb()  {
	v.X = 1002
}

func (v *Vertex) getc()  {

	fmt.Println(v)
	v.Y = 1003
}

/**
   相当于修改传进来的对象
 */
//func (v *Vertex) getby(x int) (*Vertex) {
//
//	if v.Y < 100 {
//		v.Y += x
//	}
//
//	fmt.Println("------------",100^2)
//	v.X += 函数.Pingf(x)
//	return v
//}

func adbb(v *Vertex)  {

	fmt.Println(**&v)
	v.getb()
	v.getc()
	v.Z = "4545"
}

func main() {

	//var mapj =make(map[string]Vertex)
	//mapj["asd"] = Vertex{}
	//if _, ok := mapj["as1d"]; !ok { // 判断appKey是否存在于BaseMapping中
	//	fmt.Println("exit")
	//}

	kone := make([]byte,2)

	kone[0] = -1

	fmt.Println(kone)



	//v := Vertex{1, 2, "first"}
	//
	//adbb(&v)
	//
	//fmt.Println(v)

	//fmt.Println(v)
	//fmt.Println(v.getx(100))
	//vbn := v.getby(123)
	//fmt.Println(*vbn)
	//fmt.Println(Vertex{1, 2})
	//
	//v := Vertex{1, 2}
	//v.X = 4
	//fmt.Println(v.X,v.Y)
	//
	//
	//vb := Vertex{1, 2}
	//p := &vb
	//p.X = 1e9
	//vb.X=52656
	//fmt.Println(vb)
}
