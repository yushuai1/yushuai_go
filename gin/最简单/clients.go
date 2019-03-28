package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/imroc/req"
	"time"
)


func HelpRead(resp *http.Response)  {
	byte3,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte3))
}

var Request *req.Req

func init() {
	Request = newReq()
}

func newReq() *req.Req {
	tq := &req.Req{}
	tq.SetTimeout(5 * time.Second)
	return tq
}
func main(){
	// 调用最基本的GET,并获得返回值
	t1 := time.Now().UnixNano()/1e6
	resp,err:= Request.Get("http://127.0.0.1:8825/test1")
	t2 := time.Now().UnixNano()/1e6
	if err!=nil {
		fmt.Println("-----------------",(t2-t1))
	}else {
		HelpRead(resp.Response())
	}


	//// 调用最基本的POST,并获得返回值
	//t3 := time.Now().UnixNano()/1e6
	//resp,err = http.Post("http://127.0.0.1:8888/test2", "",strings.NewReader(""))
	//t4 := time.Now().UnixNano()/1e6
	//if err!=nil {
	//	fmt.Println("-----------------",(t4-t3))
	//}else {
	//	HelpRead(resp)
	//}
}