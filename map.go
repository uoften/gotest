package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//通过v,ok := m["c"];判断键值是否存在
	//m := make(map[string]int)
	//m["a"] = 1
	//m["b"] = 11
	//asd:=m["a"]
	//asd = 12
	//fmt.Println(m)
	//fmt.Println(asd)
	//if v,ok := m["c"]; ok {
	//	fmt.Println(v)
	//}else{
	//	fmt.Println(00)
	//}

	a := new([]int)
	fmt.Println(a) //输出&[]，ａ本身是一个地址
	b := make([]int,3)
	fmt.Println(b) //输出[0]，ｂ本身是一个slice对象，其内容默认为０

	//按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	fmt.Printf("%#v\n",keys)
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
