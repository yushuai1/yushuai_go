package main

import (
	"strings"
	"net/http"
	"io/ioutil"
	"fmt"
)

func helpReads(resp *http.Response)  {
	byte3,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte3))
}

func main() {
	// 下面测试binding数据
	// 首先测试binding-JSON,
	// 注意Body中的数据必须是JSON格式
	resp, _ := http.Post("http://127.0.0.1:80/feature",
		"application/json",
		strings.NewReader("{\"user\":\"TAO\", \"password\": \"123\"}"))
	helpReads(resp)

}
