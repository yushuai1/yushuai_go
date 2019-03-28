package main

import (
	"fmt"
	_ "strconv"
	_ "time"
	"strconv"
	"time"
)


func Swap(x, y string) (string, string) {
	return y, x
}

func Pingf(x int) (int) {
	return x^2
}

func main(){
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	timestamp = timestamp[:10]
	fmt.Println(timestamp)

	sum := 0
	for {
		sum ++
		if sum > 2{
			break
		}else{
			fmt.Println(sum)
		}
	}

    result1,result2 :=Swap("yushuai","nihao")
	fmt.Println(result1, result2)

}
