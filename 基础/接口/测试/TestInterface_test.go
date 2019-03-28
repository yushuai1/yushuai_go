package yyyh

import (
	"testing"
	"fmt"
)

func TestRun(t *testing.T) {
	i:=1
	v:=i
	v = v+1
	fmt.Println(i,v)

	Run()
}