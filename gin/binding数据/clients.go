package main

import (
	"strings"
	"net/http"
	"io/ioutil"
	"fmt"
)

func helpRead(resp *http.Response)  {
	byte3,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte3))
}

func main() {
	// 下面测试binding数据
	// 首先测试binding-JSON,
	// 注意Body中的数据必须是JSON格式
	resp, _ := http.Post("http://0.0.0.0:8888/bindJSON",
		"application/json",
			strings.NewReader("{\"user\":\"TAO\", \"password\": \"123\"}"))
	helpRead(resp)

	// 下面测试bind FORM数据
	resp, _ = http.Post("http://0.0.0.0:8888/bindForm",
		"application/x-www-form-urlencoded",
			strings.NewReader("user=TAO&password=123"))
	helpRead(resp)

	// 下面测试接收JSON和XML数据
	resp, _ = http.Get("http://0.0.0.0:8888/someJSON")
	helpRead(resp)
	resp, _ = http.Get("http://0.0.0.0:8888/moreJSON")
	helpRead(resp)
	resp, _ = http.Get("http://0.0.0.0:8888/someXML")
	helpRead(resp)
}
