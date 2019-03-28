package main

import (
	"fmt"
	"time"
)

func main() {

	var jjj int
	hj := make(map[string]interface{})

	hj["y"] = 100

	jjj = hj["y"].(int)

	fmt.Println(time.Now(),jjj)
}
