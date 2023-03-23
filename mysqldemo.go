package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Gorm *gorm.DB
	Srv  http.Server
)

func main() {
	InitMysqlConn()
	InitDb()
	ctx := InitServer(context.Background())
	<-ctx.Done()

}

func InitMysqlConn() {
	dsn := "root:root@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}else{
		log.Println("mysql连接成功")
	}
}

func InitDb() {
	Gorm.AutoMigrate(&Department{})
	Gorm.AutoMigrate(&User{})

	var dep = &Department{
		Name: "TestDepart",
	}
	Gorm.Create(dep)
	Gorm.Create(&User{
		Name:       "TestName",
		Age:        20,
		Department: *dep,
	})
}

func InitServer(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	func() {
		http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				// fmt.Println(r.)
				var user User
				result := Gorm.Preload("Department").First(&user)
				if result.Error != nil {
					w.Write([]byte(result.Error.Error()))
					return
				}
				rmsg, err := json.Marshal(user)
				if err != nil {
					w.Write([]byte(err.Error()))
					return
				}
				w.Write(rmsg)
				return
			default:
				fmt.Println("Not a support requeset method.")
				return
			}
		})
	}()

	go func() {
		Srv.Addr = ":8899"
		Srv.ListenAndServe()
		cancel()
	}()

	go func() {
		fmt.Println("Service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		Srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}

type Department struct {
	Id   uint   `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name" json:"name"`
}

type User struct {
	Id           uint       `gorm:"column:id;primaryKey"`
	Name         string     `gorm:"column:name" json:"name"`
	Age          uint       `gorm:"column:age" json:"age"`
	DepartmentId uint       `gorm:"column:department_id" json:"department_id"`
	Department   Department `gorm:"forigenKey:DepartmentId" json:"department"`  // 这里可以随便写tag的位置，中间可以加空格，或者加；
}
