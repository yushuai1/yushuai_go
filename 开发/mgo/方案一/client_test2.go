package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
)

type Diary struct {
	Uid        bson.ObjectId `bson:"uid"`
	ID         bson.ObjectId `bson:"_id"`
	CreatTime  time.Time     `bson:"creattime"`
	UpdateTime time.Time     `bson:"updatetime"`
	Title      string        `bson:"title"`
	Content    string        `bson:"content"`
	Mood       int           `bson:'Mood"`
	Pic        []string      `bson:'pic'`
}

//通过uid查找本作者文章，并且显示文章作者名字
func FindDiary(uid string) ([]interface{}, error) {
	con := GetDataBase().C("diary")
	// 其中的lookup功能可以实现类似于mysql中的join操作，方便于关联查询。
	/*对应mongo命令行
	 db.diary.aggregate([{$match:{uid: ObjectId("58e7a1b89b5099fdc585d370")}},
	 {$lookup{from:"user",localField:"uid",foreignField:"_id",as:"user"}},
	 {$project:{"user.name":1,title:1,content:1,mood:1}}]).pretty()
	*/
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"uid": bson.ObjectIdHex(uid)}},
		bson.M{"$lookup": bson.M{"from": "user", "localField": "uid", "foreignField": "_id", "as": "user"}},
		bson.M{"$project": bson.M{"user.name": 1, "title": 1, "content": 1, "mood": 1, "creattime": 1}},
	}
	pipe := con.Pipe(pipeline)
	var data []interface{}
	err := pipe.All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func ModifyDiary(id, title, content string) (err error) {
	con := GetDataBase().C("diary")
	//更新
	/*对应mongo命令行
	  db.diary.update({_id:ObjectId("58e7a1b89b5099fdc585d370")},
	   {$set:{title:"modify title",content:"modify content",
	   updatetime:new Date()})*/
	err = con.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"title": title, "content": content, "updatetime": time.Now().Add(8 * time.Hour)}})
	return

}