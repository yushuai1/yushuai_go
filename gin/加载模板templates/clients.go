package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func helpRead(resp *http.Response) {
	byte3, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte3))
}
func main() {
	// 测试加载HTML模板
	resp, _ := http.Get("http://0.0.0.0:8888/index")
	helpRead(resp)
}
