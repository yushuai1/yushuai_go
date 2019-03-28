package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func HelpReads(resp *http.Response)  {
	byte3,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte3))
}

func main() {

	//port:="48888"
	//name:="tst"
	url := "http://10.2.1.227:5555/sign-in"
	resp, err :=http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("port=6666"))

	//HelpReads(resp)
	fmt.Println(resp, err)


}