package main

import (
	time "time"
	"fmt"
)
import "math/rand"

func main() {

	res:=0

	s1 := make([]byte, 512)
	for i := 0; i < 512; i++ {
		s1[i] = byte(rand.Intn(127))
	}

	s2 := make([]byte, 512)
	for i := 0; i < 512; i++ {
		s2[i] = byte(rand.Intn(127))
	}
	t1 := time.Now()
	for x := 0; x < 500000; x++ {
		for y := 0; y < 19; y++ {
			for z := 0; z < 512; z++ {
				if (s2[z] & s1[z]) == 1 {
					res++
				}
			}
		}
	}

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
