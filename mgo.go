package main

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)
var GlobalSession *mgo.Session

type Host struct {
	Id       bson.ObjectId `bson:"_id"`
}

func init() {
	// 连接数据库
	dialInfo := &mgo.DialInfo{
		Addrs: []string{"127.0.0.1:27017"}, //远程(或本地)服务器地址及端口号
		Direct: false,
		Timeout: time.Second * 3,
		Username: "admin",
		Password: "Marvelnet123",
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println("mgo链接错误")
	}
	GlobalSession = session
}
type TopologyVirtualInterface struct {
	Id                 bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	NextDeviceIp       string        `json:"next_device_ip" bson:"next_device_ip"`
	DeviceName         string        `json:"device_name" bson:"device_name"`
	InterfaceName      string        `json:"interface_name"  bson:"interface_name"`
	InterfaceAliasName string        `json:"interface_alias_name"  bson:"interface_alias_name"`
	InterfaceIp        string        `json:"interface_ip" bson:"interface_ip"`
	UpdateTime         string        `json:"update_time"  bson:"update_time"`
	Uuid               string        `json:"uuid"  bson:"uuid"`
	NextIp             string        `json:"next_ip" bson:"next_ip"`
	NextInterfaceName  string        `json:"next_interface_name"  bson:"next_interface_name"`

}
//公共model方法
func withCollection(tableName string , f func(*mgo.Collection) error) error {
	s := GlobalSession
	session := s.Copy()
	defer session.Close()
	return f(session.DB("nacs").C(tableName))
}
func Count(tableName string, selector interface{}) (count int, err error) {
	err = withCollection(tableName, func(c *mgo.Collection) error {
		count, err = c.Find(selector).Count()
		if err != nil {
			return err
		} else {
			return nil
		}
	})
	return
}
func FindOne(tableName string, selector interface{}) (retData TopologyVirtualInterface, err error) {
	err = withCollection(tableName, func(c *mgo.Collection) error {
		return c.Find(selector).One(&retData)
	})
	return
}

func Find(tableName string, selector interface{}) (retData []interface{}, err error) {
	err = withCollection(tableName, func(c *mgo.Collection) error {
		return c.Find(selector).All(&retData)
	})
	return
}

func main()  {
	//获取用户输入
	var t string
	flag.StringVar(&t, "t", "topology_virtual_interface", "表名")
	flag.Parse()
	if t == "" {
		fmt.Println("缺少参数：-t=表名")
		return
	}
	//查询表记录数
	routeCount,err := Count(t,bson.M{})
	if err != nil {
		fmt.Println("查询表记录数错误")
		return
	}
	fmt.Println("表"+t+"记录总数：")
	fmt.Println(routeCount)

	//查询多条记录
	//用routeData==nil判断空值
	routeData,err := Find(t,bson.M{})
	if err != nil {
		fmt.Printf("查询表记录数错误:%v",err)
		return
	}
	fmt.Println(routeData==nil)
	fmt.Printf("%#v\n",routeData)

	//查询单条记录
	//用err=nil判断空值
	retData,err := FindOne(t,bson.M{})
	if err != nil {
		fmt.Printf("查询表记录数错误:%v",err)
		return
	}
	fmt.Printf("%#v\n",retData)
}