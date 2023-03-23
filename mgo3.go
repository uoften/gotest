package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	DeviceOrderDbName         = "sun9ShouMi_device_order"
	DeviceOrderCollectionName = "device_order"
)
type DeviceOrderInfo struct {
	Id           bson.ObjectId        `bson:"_id"`
	AgentId      string               `bson:"agent_id"`
	AgentPhone   string               `bson:"agent_phone"`
	DeviceSn     string               `bson:"device_sn"`
	DeviceType   int32                `bson:"device_type"`
	DeviceCounts int32                `bson:"device_counts"`
	CreatedAt    bson.MongoTimestamp  `bson:"created_at"`
	Status       int32                `bson:"status"`
	Distatus     int32                `bson:"distatus"`
	MchtId       string               `bson:"mcht_id"`
	MchtName     string               `bson:"mcht_name"`
	MchtPhone    string               `bson:"mcht_phone"`
	StoreName    string               `bson:"store_name"`
}

func main() {
	// 连接数据库
	session, err := mgo.Dial("192.168.10.198:27018")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 获取数据库,获取集合
	c := session.DB(DeviceOrderDbName).C(DeviceOrderCollectionName)

	// 存储数据
	m1 := DeviceOrderInfo{bson.NewObjectId(), "5f6ad0c6e138236f9ca5a1b8","17300000000","SN8888888",2,200,bson.MongoTimestamp(time.Now().Unix()<<32 + 1),3,1,"5f6ad0c6e138236f9ca5a1b8","","17300000000","手信商城"}
	//m2 := DeviceOrderInfo{bson.NewObjectId(), "user1", 11, "user2@dotcoo.com"}
	//m3 := DeviceOrderInfo{bson.NewObjectId(), "user3", 11, "user3@dotcoo.com"}
	//m4 := DeviceOrderInfo{bson.NewObjectId(), "user3", 11, "user4@dotcoo.com"}
	//err = c.Insert(&m1, &m2, &m3, &m4)
	err = c.Insert(&m1)
	if err != nil {
		panic(err)
	}

	// 读取数据
	var ms = []DeviceOrderInfo{}
	err = c.Find(&bson.M{"AgentId":"5f6ad09ee138236f9ca5a1b7"}).All(&ms)
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
