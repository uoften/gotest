package main

import (
	"fmt"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

type Xi []int
type Xs []string

func (p Xi) Len() int               { return len(p) }
func (p Xi) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func (p Xs) Len() int               { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

// 用接口实现好了排序方法，但是我不知道你会传入什么进行排序。于是我写了个接口，你只要实现了接口的全部方法，那么我就允许你调用我的服务，进行排序。
func main() {
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	ints = append(ints,99)
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}
	fmt.Printf("%#v\n",ints)
	fmt.Printf("%#v\n",strings)
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)
}