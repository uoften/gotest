package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	id     int    //	首字母小写、无 tag
	name   string `json:"name"` //	首字母小写、有 tag
	Age    int    //	首字母大写、无 tag
	Mobile string `json:"mobile"` // 首字母大写、有 tag
}

func main() {
	s := Student{
		id:     20210901,
		name:   "wohu",
		Age:    20,
		Mobile: "123456",
	}
	fmt.Printf("转换 json 前的内容为 %+v\n", s) // %+v类似%v，但输出结构体时会添加字段名
	jsonContent, _ := json.MarshalIndent(s, "", "    ")
	fmt.Printf("转换 json 后的内容为\n%v\n", string(jsonContent))
}
//结构体成员首字母小写，无论是否加 tag 都无法转换为 json 字段，即 json 中会丢弃首字母小写的字段值；
//结构体成员首字母大写，分以下两种情况：
//不加 tag 转换为 json 后的字段名和结构体当前的字段名一致；
//加 tag 转换为 json 后的字段名与 tag 里面的字段名一致，结构体中原来的值就被抛弃；