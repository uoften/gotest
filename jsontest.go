package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string `json:"-"` //此字段不会输出到屏幕
	//Company  string   `json:"company"`       这样打印输出别名首字母就会小写(二次编码)
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:",string"`
	Price    float64  `json:",string"`
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"Golang", "PHP", "Java", "C++"}, true, 666.666}

	//编码，根据内容生成json文本
	//buf, err := json.Marshal(s)
	//buf =  {"subjects":["Golang","PHP","Java","C++"],"IsOk":"true","Price":"666.666"}
	buf, err := json.MarshalIndent(s, "", "    ") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))

	//s := []map[string]interface{}{}
	//m1 := map[string]interface{}{"name": "John", "age": 10}
	//m2 := map[string]interface{}{"name": "Alex", "age": 12}
	//
	//s = append(s, m1, m2)
	//s = append(s, m2)
	//
	//b, err := json.Marshal(s)
	//if err != nil {
	//	fmt.Println("json.Marshal failed:", err)
	//	return
	//}
	//
	//fmt.Println(string(b))

	//m := make(map[string]interface{}, 4) //因为类型多，可以用interface空接口
	//m["address"] = "北京"
	//m["languages"] = []string{"Golang", "PHP", "Java", "Python"}
	//m["status"] = true
	//m["price"] = 666.666
	//
	////编码成json
	////result, err := json.Marshal(m)
	////result =  {"address":"北京","languages":["Golang","PHP","Java","Python"],"price":666.666,"status":true}
	//result, err := json.MarshalIndent(m, "", "	")
	//if err != nil {
	//	fmt.Println("err = ", err)
	//	return
	//}
	//fmt.Println(string(result))

}