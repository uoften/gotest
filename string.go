package main

import (
	"fmt"
	"strings"
)

func main() {

	//方法1
	s1 := "abcdefgabc"
	s2 := []byte(s1)
	s2[1] = 'B'
	fmt.Println(string(s2)) //aBcdefgabc
	//方法2
	s3 := []rune(s1)
	s3[1] = 'B'
	fmt.Println(string(s3)) //aBcdefgabc
	//方法3
	newStr := "ABC"
	old := "abc"
	s4 := strings.Replace(s1, old, newStr, 2)
	fmt.Println(s4) //ABCdefgABC
}