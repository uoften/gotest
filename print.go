package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3}
	fmt.Println(s1, "哈哈") //[1 2 3]

	s2 := s1
	fmt.Println(s1, "哈哈") //[1 2 3]
	for i := 0; i < 3; i++ {
		s2[i] = s2[i] + 1
	}
	fmt.Println(s1) //[2 3 4]
	fmt.Println(s2) //[2 3 4]
	var a int = 10
	var b int = 20
	a=a+b
	b=a-b
	a=a-b
	fmt.Printf("a = %v\nb = %v\n",a,b)
}