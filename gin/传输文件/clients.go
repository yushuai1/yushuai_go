package main

import (
	"bytes"
	"mime/multipart"
	"os"
	"io"
	"net/http"
	"io/ioutil"
	"fmt"
)


func helpRead(resp *http.Response)  {
	byte3,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte3))
}


func main() {
	// 上传文件POST
	// 下面构造一个文件buf作为POST的BODY
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	fw, _ := w.CreateFormFile("uploadFile", "images.png") //这里的uploadFile必须和服务器端的FormFile-name一致
	fd, _ := os.Open("images.png")
	defer fd.Close()
	io.Copy(fw, fd)
	w.Close()
	resp, _ := http.Post("http://0.0.0.0:8888/upload", w.FormDataContentType(), buf)
	helpRead(resp)
}
