package main

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

var session *mgo.Session
var database *mgo.Database

func init() {
	/*配置mongodb的josn文件,配置内容如下:
	 {
	   "hosts": "localhost",
	   "database": "user"
	  }*/
	var err error

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"127.0.0.1:27017"},
		Direct:    false,
		Timeout:   time.Second * 1,
		PoolLimit: 4096, // Session.SetPoolLimit //
	}
	//创建一个维护套接字池的session
	session, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Println(err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	//使用指定数据库
	database = session.DB("test")
}

func GetMgo() *mgo.Session {
	return session
}

func GetDataBase() *mgo.Database {
	return database
}

func GetErrNotFound() error{
  return mgo.ErrNotFound
}

