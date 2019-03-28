package main

import (
	"fmt"
	"bytes"
)

func addAllByte(arg ...[]byte) []byte  {
    var jk []byte
	return bytes.Join(arg,jk)
}

func main() {
	var by1 = []byte {3,4}
	var by2 = []byte {4,6}
	var by3 = []byte {7,9}

	byteArray := make([][]byte,3)
	byteArray[0] = by1
	byteArray[1] = by2
	byteArray[2] = by3
	yu := bytes.Join(byteArray,[]byte(""))

	fmt.Println(yu)

}
