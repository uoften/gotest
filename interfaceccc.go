package main

import "fmt"

type A interface {
	ShowA() int
	ShowB() int
}
type B interface {
}

func Show(a A) int {
	return a.ShowA()+a.ShowB()
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	a := Work{3}
	fmt.Println(Show(a))
	var aa A = a
	var bb B = a
	fmt.Println(aa==bb)
	switch bb.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
//结构体b（Work）实现了接口a的ShowA，ShowB方法，所以结构体b可以调用接口a的方法Show
