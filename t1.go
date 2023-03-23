package main

import "fmt"

type Supermen struct {
	Name string
	Num int
}

func (p Supermen) test(n int) {
	fmt.Println(p.Num+n)
}

func main() {
	//结构体:属性，方法，类似php的class
	//var aa Supermen
	//aa.Num = 100
	//aa.Name = "tiom"
	//aa.test(1)
	//fmt.Println(aa.Name)

	//切片:初始化，增加元素，复制，类似php的遍历重构数组
	var aa []int = []int{
		1,2,3,4,5,6,7,8,9,
	}
	var bb []int
	var cc = make([]int,0)
	var dd = make([]int,14)
	fmt.Println(aa)
	fmt.Println(bb)
	cc = append(cc,2,3,4,5,6,7,8)
	fmt.Println(cc)
	copy(dd,cc)
	fmt.Println(dd)
}
