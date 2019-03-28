package main

import (
	"os"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"strings"
	"path/filepath"
	"io/ioutil"
)

func main() {
	var codePath string

	fmt.Scanf("%s", &codePath)

	fmt.Println(codePath,"---------")
	if f, err := os.Stat(codePath); err == nil {
		if f.IsDir() == false { // 判断是否是文件夹
			log.Error("请正确传入路径信息，路径必须是文件夹")
			return
		}
	} else {
		log.Error("请正确传入路径信息!   ", err)
		return
	}

	fmt.Println(codePath)

	x := CountDirLines(codePath, "java")
	fmt.Println(x)

}

/*
遍历目录，找出后缀符合的文件
 */
func CountDirLines(dirPth, suffix string) int64 {
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	var lineSum int64 = 0
	filepath.Walk(dirPth, func(fileName string, fi os.FileInfo, err error) error { //遍历目录

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			if l, err := CountFileLine(fileName); err == nil {
				fmt.Println(fileName)
				lineSum += l
			}
		}
		return nil
	})

	return lineSum
}

/*
统计单个文件的行数
 */
func CountFileLine(name string) (count int64, err error) {

	data, err := ioutil.ReadFile(name) // 读取文件内容

	if err != nil {
		return
	}
	count = 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	return
}
