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
	// 下面测试使用中间件
	resp, _ := http.Get("http://0.0.0.0:8888/logger")
	helpRead(resp)

	// 测试验证权限中间件BasicAuth
	resp, _ = http.Get("http://0.0.0.0:8888/admin/secrets")
	helpRead(resp)
}
