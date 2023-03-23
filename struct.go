package main

import (
	"fmt"
	"strconv"
)

type device struct {
	ip string
}

type Person2 struct {
	age int
}

func test2(d map[string]device) {
	//d.ip = "123"
	for k,_:=range d{
		var dd device
		dd.ip = "aasd"
		d[k] = dd
	}
}

func main() {
	//结构体空判断
	var st Person2
	if (Person2{} == st) {
		fmt.Println("empty")
	} else {
		fmt.Println("not empty")
	}

	d:= make(map[string]device)
	for i:=1;i<=10;i++{
		var dd device
		dd.ip = strconv.Itoa(i)
		d[strconv.Itoa(i)] = dd
	}
	fmt.Println(d)
	//引用类型传值
	test2(d)
	fmt.Println(d)

	//作用域
	main2()
	main3()
}

func main2() {
	x := 1
	println(x) // 1
	if true{
		//新的 x 变量的作用域只在代码块内部
		println(x) // 1
		x := 2
		println(x) // 2
	}
	println(x) // 1
}

func change(arr []int) {
	arr[0] = 7
	fmt.Println(arr) // [7 2 3]
}
func change2(arr [3]int) {
	arr[0] = 7
	fmt.Println(arr) // [7 2 3]
}
func main3() {
	x := [3]int{1,2,3}
	change2(x)
	fmt.Println(x) // [1 2 3] // 并不是你以为的 [7 2 3]
	fmt.Printf("指针地址为：%p\n",&x)

	//切片：前包后不包，[0:2]为所用0,1, 不包含索引2
	xx := x[0:2]
	change(xx)
	fmt.Println(xx)	// [7 2 3] // 切片是引用类型，使用指针传值
	fmt.Printf("指针地址为：%p\n",xx)
}