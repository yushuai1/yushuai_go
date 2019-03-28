package main

import (
	"fmt"
)

func powers(n,m uint64) uint64{
	if m==1 {
		return n
	}
	return n*powers(n,m-1)
}

func main() {

	var i,a,b,c uint64

	for i=3;i<40;i++  {
		for a=1; a<40;a++  {
			for b=1;b<40 ;b++  {
				for c=1;c<40 ;c++  {
					v1:= powers(a,i)
					v2:=powers(b,i)
					v3:=powers(c,i)
					if (v1+v2) ==v3 {
						fmt.Println(a,b,c,i,v1,v2,v3)
					}
				}

			}
		}
	}
}
