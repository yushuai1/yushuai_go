package main

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"

)

type User struct {
	ID       bson.ObjectId `bson:"_id"`
	UserName string        `bson:"username"`
	Summary  string        `bson:"summary"`
	Age      int           `bson:"age"`
	Phone    int           `bson:"phone"`
	PassWord string        `bson:"password"`
	Sex      int           `bson:"sex"`
	Name     string        `bson:"name"`
	Email    string        `bson:"email"`
}

func Register(password string, username string) (err error) {
	con := GetDataBase().C("user")
	//可以添加一个或多个文档
	/* 对应mongo命令行
	   db.user.insert({username:"13888888888",summary:"code",
	   age:20,phone:"13888888888"})*/
	err = con.Insert(&User{ID: bson.NewObjectId(), UserName: username, PassWord: password})
	return
}

func FindUser(username string) (User, error) {
	var user User
	con := GetDataBase().C("user")
	//通过bson.M(是一个map[string]interface{}类型)进行
	//条件筛选，达到文档查询的目的
	/* 对应mongo命令行
	  db.user.find({username:"13888888888"})*/
	if err := con.Find(bson.M{"username": username}).One(&user); err != nil {
		if err.Error() != GetErrNotFound().Error() {
			return user, err
		}

	}
	return user, nil
}

func main() {
	err :=Register("asd","asd")
	if err!=nil {
		fmt.Println(err)
	}
}