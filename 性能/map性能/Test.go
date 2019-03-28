package main

import (
	"fmt"
	"time"
	"strconv"
)

var yus,yud string

func main() {

    k:=0

	mapTest := make(map[string]string)




	t1 := time.Now().UnixNano()/1e6

	for i:=0;i<10000000;i++ {
		mapTest["v-"+strconv.Itoa(i)] = "v-"+strconv.Itoa(i)
	}


	t2 := time.Now().UnixNano()/1e6
	var s string
	for i:=0;i<100000000;i++ {
		s=mapTest["v-100"]
	}
	t8 := time.Now().UnixNano()/1e6
	fmt.Println("---------************--------",t2-t1,"---",s,"---",t8-t2)

	for _,v := range mapTest{

		k++

		yud = v

	}

	t3 := time.Now().UnixNano()/1e6

	fmt.Println(k)

	fmt.Printf("%s \n",yud)

	fmt.Println((t2-t1),(t3-t2))

	arrays := make([]string,100)

	t4:= time.Now().UnixNano()/1e6

	for n:=0;n<100;n++ {

		arrays[n] = strconv.Itoa(n)

	}

	t5 := time.Now().UnixNano()/1e6

	for n:=0;n<100;n++ {

		yud = arrays[n]

	}

	t6 := time.Now().UnixNano()/1e6

	fmt.Printf("----------------------\n")

	fmt.Println((t5-t4),(t6-t5),yud)

	fmt.Println("********************************************")

	zhizhen := make([]*[]int,2)

	sz1:=make([]int,2)

	sz2:=make([]int,2)

	zhizhen[0] = &sz1

	zhizhen[1] = &sz2

	fmt.Println(len(zhizhen),(*zhizhen[0])[0],&zhizhen[0])

	*zhizhen[0] = nil

	fmt.Println(len(zhizhen),(zhizhen[0]),sz1)

}
