package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var session *mgo.Session


func init() {

	var err error

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"127.0.0.1:27017"},
		Direct:    false,  //直接连接轮询调用
		Timeout:   time.Second * 1,
		PoolLimit: 4096, // Session.SetPoolLimit

		FailFast: true,   //快速失败
		//Database: "数据库"
		//ReplicaSetName:"副本集名称"
		//Source:"admin"   //验证权限集合
		//Service:"mongodb"   //验证权限服务名称
		//Username string
		//Password string

	}

	session, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Println(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)

}

type SessionStore struct {
	session *mgo.Session
}

//获取数据库的collection
func (d * SessionStore) C(name string) *mgo.Collection {
	return d.session.DB("test").C(name)
}

//为每一HTTP请求创建新的DataStore对象
func NewSessionStore() * SessionStore {
	ds := & SessionStore{
		session: session.Copy(),
	}
	return ds
}

func (d * SessionStore) Close() {
	d.session.Close()
}

func GetErrNotFound() error {
	return mgo.ErrNotFound
}