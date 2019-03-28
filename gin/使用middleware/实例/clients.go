package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var resp *http.Response
	resp, _ = http.Get("http://127.0.0.1:8989/ff/shaui?username=ft&&password=123")
	// resp, _ = http.Get("http://10.0.203.92:8989?username=ft&&password=123")

	helpRead(resp)
}

func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}