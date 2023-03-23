package main

import "fmt"

func main() {
	var ret []byte
	ret = []byte{32,10,7,27}
	fmt.Println(ret)
	fmt.Println(string(ret))
}