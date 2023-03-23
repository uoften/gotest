package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	fmt.Println(a)
	b := s[:2]
	fmt.Println(b)
	c := s[1:2:cap(s)]
	fmt.Println(cap(c))
}
