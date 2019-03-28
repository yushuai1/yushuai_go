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
	// 下面测试重定向
	resp, _ := http.Get("http://0.0.0.0:8888/redirect")
	helpRead(resp)
}
