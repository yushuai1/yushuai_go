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

	var yu int32 = 100

	gh := fmt.Sprintf("%d",yu)

	fmt.Println(gh)

	// 使用post_form形式,注意必须要设置Post的type,同时此方法中忽略URL中带的参数,所有的参数需要从Body中获得
	//resp, _ := http.Post("http://localhost:58888/interface/iris/register",
	//	"application/x-www-form-urlencoded",
	//	strings.NewReader("?custCode=8888888&base64Image=yushuai"))
	resp, _ := http.Post("http://localhost:8877/match",
		"application/x-www-form-urlencoded",
		strings.NewReader("pic=8888888"))

	helpRead(resp)

}
