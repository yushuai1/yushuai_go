package main

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id_       bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Age       int           `bson:"age"`
	arr []string      `bson:"arr"`
}

func main() {
	fmt.Println("This is a test to use mgo for go.")

	//connect server
	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect success.")
	}
	defer session.Close()


	//设置集群模式
	session.SetMode(mgo.Monotonic, true)
	//切换数据库
	db := session.DB("test")

	//切换集合
	c := db.C("people")

	//插入数据
	err = c.Insert(&User{
		Id_:       bson.NewObjectId(),
		Name:      "Jimmy Kuu",
		Age:       33,
		arr: []string{"Develop", "Movie"},
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert success.")
	}

	//查找所有数据
	var users []User
	c.Find(nil).All(&users)
	fmt.Println(users)

	//根据条件查找
	c.Find(bson.M{"name": "Jimmy Kuu"}).All(&users)
	fmt.Println(users)

	//数据修改
	c.Upsert(bson.M{"_id": bson.ObjectIdHex("5a911d109c44bc1a30c9472d")},
		bson.M{"$set": bson.M{
			"name": "Jimmy Gu",
			"age":  34,
		}})
	id := "5a911f559c44bc07a4fc612a"
	objectId := bson.ObjectIdHex(id)
	user := new(User)
	c.Find(bson.M{"_id": objectId}).One(&user)
	fmt.Println(user)

	//update add
	c.Update(bson.M{"_id": bson.ObjectIdHex("5a911f559c44bc07a4fc612a")},
		bson.M{"$inc": bson.M{
			"age": -1,
		}})
	objectId = bson.ObjectIdHex(id)
	c.FindId(objectId).One(&user)
	fmt.Println(user)

	//add ele
	c.Update(bson.M{"_id": bson.ObjectIdHex("5a911f559c44bc07a4fc612a")},
		bson.M{"$push": bson.M{
			"interests": "Golang",
		}})
	objectId = bson.ObjectIdHex(id)
	c.FindId(objectId).One(&user)
	fmt.Println(user)

	//del ele
	c.Update(bson.M{"_id": bson.ObjectIdHex("5a911f559c44bc07a4fc612a")},
		bson.M{"$pull": bson.M{
			"interests": "Golang",
		}})
	objectId = bson.ObjectIdHex(id)
	c.FindId(objectId).One(&user)
	fmt.Println(user)

	//remove
	c.Remove(bson.M{"name": "Jimmy Kuu"})
	c.Find(nil).All(&users)
	fmt.Println(users)
}
