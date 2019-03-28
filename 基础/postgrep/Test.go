package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"

	"time"
)

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


var connection gorose.Connection
var err error

func init()  {
	connection, err = gorose.Open(dbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	//connection, err := gorose.Open(config.DbConfig, "mysql_dev")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	// close DB
	defer connection.Close()

	for k := 0; k < 5; k++ {
		go transTest(k)
	}
	time.Sleep(time.Second*1)
}

func transTest(k int)  {
	//var data datas


	db := connection.GetInstance()

	trans,err := db.Transaction(func() error {
		jiev := make(map[string]interface{})
		jiev["classes"] = "classes"
		res33,err31 := db.Table("timees").Data(jiev).Insert()
		if err31!=nil||res33<1 {
			fmt.Println(err31)
			db.Rollback()
		}

		return nil
	})

	fmt.Println(trans, err)
}