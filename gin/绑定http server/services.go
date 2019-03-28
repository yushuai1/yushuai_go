package main

import (
	"time"
	"net/http"
)

func main() {

	// 方法二
	http.ListenAndServe(":8888", router)

	// 方法三:
	server := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}
