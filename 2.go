package main

import (
	"fmt"
	"net"
)

import "github.com/go-sql-driver/mysql” // 具体的驱动包
import "database/sql"

// 初始化连接
func initDB() (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	// todo 不要在这里关闭它, 函数一结束,defer就执行了
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}