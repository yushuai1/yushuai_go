package main

import (
	"fmt"
)



func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar(p *int) {
	fmt.Println(*p)
}

func main() {
	var p *[]string
	//p, err := foo()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//bar(p)
	fmt.Println(*p)
}