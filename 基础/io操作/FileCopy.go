package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	copylen, err := copyFile("dst.txt", "src.txt")
	if err != nil {
		return
	} else {
		fmt.Println(copylen)
	}

}

//函数copyFile的功能是将源文件sec的数据复制给dst
func copyFile(dstName, srcName string) (copylen int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	//当return时就会调用src.Close()把源文件关闭
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	//当return是就会调用src.Close()把目标文件关闭
	defer dst.Close()
	return io.Copy(dst, src)
}