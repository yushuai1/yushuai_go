package main

import "fmt"

func main() {


	a := make([]string,0)

	a = append(a,"ui")
	fmt.Println(&a)
	fmt.Sprintf("%p",&a)

	b:=make([][]string,0)
	b = append(b,a)


	b[0] = append(b[0],"shuai")

	fmt.Println(b[0])
}
