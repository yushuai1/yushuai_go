package main

import "github.com/gohouse/gorose"
import "fmt"
//import _"github.com/lib/pq"
import (
	//"github.com/goinggo/mapstructure"
	//"com.irisking/models/person"
	//"reflect"
	//"com.irisking/util/db"
	_ "github.com/go-sql-driver/mysql"
	//"time"
	//"com.irisking/util/db"
	"sync"
	//"time"
	//"github.com/henrylee2cn/go-logging/color"
	//"crypto/aes"
	"time"
)
//
//var dbConfig = map[string]map[string]string {
//	"mysql": {
//		"host":     "localhost",
//		"username": "root",
//		"password": "root",
//		"port":     "3306",
//		"database": "test",
//		"charset":  "utf8",
//		"protocol": "tcp",
//	},
//	"progrep": {
//		"host":     "localhost",
//		"username": "root",
//		"password": "",
//		"port":     "3306",
//		"database": "gorose",
//		"charset":  "utf8",
//		"protocol": "tcp",
//	},
//}

var dbConfig = map[string]interface{}{
	// 默认的链接数据库配置
	"Default": "mysql",
	// 最大连接数
	"SetMaxOpenConns": 300,
	// 最大空闲连接数
	"SetMaxIdleConns": 10,
	// 配置
	"Connections":map[string]map[string]string{
		"postgres": map[string]string{
			"host":     "10.2.1.246",
			"username": "postgres",
			"password": "ok",
			"port":     "5432",
			"database": "micro_cloud",
			"charset":  "utf8",
			"protocol": "tcp",
			"prefix":   "",      // Table prefix
			"driver":   "postgres", // Database driver(mysql,sqlite,postgres,oracle,mssql)
		},
		"mysql": map[string]string{
			"host":     "127.0.0.1",
			"username": "root",
			"password": "root",
			"port":     "3306",
			"database": "test",
			"charset":  "utf8",
			"protocol": "tcp",
			"prefix":   "",      // Table prefix
			"driver":   "mysql", // Database driver(mysql,sqlite,postgres,oracle,mssql)
		},
	},
}

type Person struct {
	Id int32 `sql:"id"`
	Adder string `sql:"adder"`
	Classes string `sql:"classes"`
	Phone string  `sql:"phone"`
}

type Testyu struct {
    Id int32
	Name string
	Feature []byte
	Test []byte
}

var con gorose.Connection
var once sync.Once
func getcon()(*gorose.Connection) {
	once.Do(func() {
		var err error
		// 加载database
		con, _ = gorose.Open(dbConfig)

		if err != nil {
			fmt.Println(err)
			panic("数据库链接失败")
		}
		errs := con.Ping()
		if errs != nil {
			fmt.Println(err)
			panic("数据库链接失败")
		}

	})
    return &con
}

func Test1()  {
	bn :=getcon()
	db := bn.GetInstance()
	db.Begin()
	jiev := make(map[string]interface{})
	jiev["classes"] = "classes"
	_,err31 := db.Table("timees").Data(jiev).Insert()

	//_,err31 := db.Query("insert into timees(classes) values('man')")
	if err31!=nil{

		fmt.Println(err31)
		db.Rollback()
		return
	}
	db.Commit()
	fmt.Println("-----")
}

func main() {

	for k:=0;k<70;k++ {
		 go Test1()
	}

	time.Sleep(time.Second*10)


	//jiev := make(map[string]interface{})
	//jiev["id"] = 102
	//jiev["classes"] = "classes"
	//db.Begin()
	//res3,err1 := db.Table("user_test").Data(jie).Insert()
	//fmt.Println(res3,err1)

	//res33,err31 := db.Table("timees").Data(jiev).Insert()
	//fmt.Println(res33,err31)
	//db.Commit()
	//db.Rollback()
	//
	//bns :=getcon()
	//dbs := bns.GetInstance()
	//t1 := time.Now().UnixNano()/1e6
	//for i:=0;i<10000 ;i++  {
	//	res, errs := dbs.Table("user_test").Get()
	//	if errs!=nil {
	//		fmt.Println(res)
     //       break
	//	}
	//}
	//t2 := time.Now().UnixNano()/1e6
	//fmt.Println(t2-t1)



	//res1,err := db.Table("testyu").Where("id=9").First()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(db.LastSql)

	//fmt.Println(res["classes"])

	//res1,_ := db.Query("select * from testyu where id=19")

	//fmt.Println(res1["feature"],"***************************************************************")
	//t := reflect.TypeOf(res1["feature"])         //反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
	//fmt.Println("Type:", t.Name()) //调用t.Name方法来获取这个类型的名称
	//var testyu Testyu
	//
	////将 map 转换为指定的结构体
	//if err := mapstructure.Decode(res1, &testyu); err != nil {
		//fmt.Println(err)
	//}
	//
    //fmt.Println(testyu,"------------",testyu.Feature)
	//
	//
    //jie := make(map[string]interface{})
    //jie["name"] = "于帅"
	//jie["feature"] = []byte{1,2,3,4,5,6}
	//jie["test"] = []byte{1,2,3,4,5,6,45,55,44}
	//res3,err1 := db.Table("testyu").Data(jie).Insert()
	//if err1!=nil {
		//fmt.Println(err1)
	//}
	//
    //fmt.Println(res3)

}

