package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	c, err := redis.Dial("tcp", "10.2.1.35:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	var username string
	var errs error
	t1 := time.Now().UnixNano()/1e6

	for ivssa:=0;ivssa<100000 ;ivssa++  {
		username, errs = redis.String(c.Do("GET", "yu"))
	}
	t2 := time.Now().UnixNano()/1e6
	fmt.Println(t2-t1)
	if errs != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	defer c.Close()
}
