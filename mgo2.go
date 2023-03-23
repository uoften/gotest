package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mail struct {
	Id bson.ObjectId "_id"
	UserRoleId   string           	 "user_role_id"
}

func main() {
	// 连接数据库
	session, err := mgo.Dial("192.168.10.198:27018")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 获取数据库,获取集合
	c := session.DB("sun9ShouMi_basic").C("app_user_info")

	// 存储数据
	//m1 := Mail{bson.NewObjectId(), "5f5097efd41c0305400c8efe"}
	//m2 := Mail{bson.NewObjectId(), "user1", 11, "user2@dotcoo.com"}
	//m3 := Mail{bson.NewObjectId(), "user3", 11, "user3@dotcoo.com"}
	//m4 := Mail{bson.NewObjectId(), "user3", 11, "user4@dotcoo.com"}
	//err = c.Insert(&m1, &m2, &m3, &m4)
	//err = c.Insert(&m1)
	//if err != nil {
	//	panic(err)
	//}

	// 读取数据
	var ms = []Mail{}
	err = c.Find(&bson.M{"name":"HM73580923"}).All(&ms)
	if err != nil {
		panic(err)
	}
	js,err := json.Marshal(ms)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
	fmt.Println(string(js))
	// 显示数据
	//for i, m := range ms {
	//	fmt.Printf("%d, %s, %d, %s\n", i, m.Id.Hex(),m.Age, m.Email)
	//}
}
