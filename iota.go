package main

import "fmt"

const (
	v = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main()  {
	fmt.Println(v,y,z,k,p)
}
