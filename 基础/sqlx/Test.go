package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //初始化一个mysql驱动，必须
	"github.com/jmoiron/sqlx"
	"time"
)



var Db *sqlx.DB

func init() {
	//"mysql"指定数据库类型，  /test指定打开的数据库  root:123 冒号隔开密码 root账号 123密码
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func run()  {
	//事务操作
	conn, err := Db.Begin()
	if err != nil {
		return
	}
	//插入数据入表
	f, err := Db.Exec("insert into timees(classes)values(?)", "man")
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	conn.Commit()
	fmt.Println(f)
}
func main() {

	for yu:=0;yu<10 ;yu++  {
		go run()
	}

	time.Sleep(time.Second*10)
}