package main

import (

	"gopkg.in/mgo.v2/bson"
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

func FindUser(username string) (User, error) {
	var user User
	ds := NewSessionStore()
	defer ds.Close()
	con := ds.C("user")
	if err := con.Find(bson.M{"username": username}).One(&user); err != nil {
		if err.Error() != GetErrNotFound().Error() {
			return user, err
		}

	}
	return user, nil
}
