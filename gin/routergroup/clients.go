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
	// 下面测试router 的GROUP
	resp, _ := http.Get("http://0.0.0.0:8888/g1/read1")
	helpRead(resp)
	resp, _ = http.Get("http://0.0.0.0:8888/g1/read2")
	helpRead(resp)
	resp, _ = http.Post("http://0.0.0.0:8888/g2/write1", "", strings.NewReader(""))
	helpRead(resp)
	resp, _ = http.Post("http://0.0.0.0:8888/g2/write2", "", strings.NewReader(""))
	helpRead(resp)
}
