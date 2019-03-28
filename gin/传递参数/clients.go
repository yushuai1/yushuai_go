package main

import (
	//"strings"
	"net/http"
	"io/ioutil"
	"fmt"
)
func helpRead(resp *http.Response)  {
	byte3,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte3))
}
func main() {
	// GET传参数,使用gin的Param解析格式: /test3/:name/:passwd
	resp, _ := http.Get("http://0.0.0.0:8888/test3/yu=45-47-41/shuai=4564")
	helpRead(resp)

	//// POST传参数,使用gin的Param解析格式: /test3/:name/:passwd
	//resp, _ = http.Post("http://0.0.0.0:8888/test4/name=PT/passwd=456", "", strings.NewReader(""))
	//helpRead(resp)
	//
	//// 注意Param中':'和'*'的区别
	resp, _ = http.Get("http://0.0.0.0:8888/test5/name=TAO/passwd=789")
	helpRead(resp)
	//resp, _ = http.Get("http://0.0.0.0:8888/test5/name=TAO/")
	//helpRead(resp)
}
