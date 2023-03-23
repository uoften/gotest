package main

import (
	"fmt"
	"hash/fnv"
)

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	fmt.Println(Hash("a"))
}

